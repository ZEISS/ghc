package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zeiss/ghc/internal/cfg"
	"github.com/zeiss/ghc/pkg/spec"

	"github.com/spf13/cobra"
)

var config = cfg.Default()

const (
	versionFmt = "%s (%s %s)"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(ValidateCmd)
	RootCmd.AddCommand(InstallCmd)

	RootCmd.Flags().StringVarP(&config.Root.Run, "run", "r", config.Root.Run, "run a specific hook")

	RootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", config.Verbose, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&config.Force, "force", "f", config.Force, "force overwrite")
	RootCmd.PersistentFlags().StringVarP(&config.File, "config", "c", config.File, "config file")

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true
}

var RootCmd = &cobra.Command{
	Use:   "ghc",
	Short: "ghc is a teeny tiny tool to manage git hooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
}

func runRoot(ctx context.Context) error {
	cfg := filepath.Clean(config.File)
	cwd := filepath.Dir(cfg)

	s, err := os.ReadFile(cfg)
	if err != nil {
		return err
	}

	var spec spec.Spec
	if err := spec.UnmarshalYAML(s); err != nil {
		return err
	}

	cmds, err := spec.Hook(config.Root.Run)
	if err != nil {
		return err
	}

	for _, c := range cmds {
		cc := strings.Split(c, " ")

		c := exec.CommandContext(ctx, cc[0], cc[1:]...)
		c.Dir = cwd
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			return err
		}
	}

	return nil
}
