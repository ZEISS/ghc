package cmd

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zeiss/ghc/internal/cfg"
	"github.com/zeiss/ghc/pkg/spec"

	"github.com/spf13/cobra"
)

var config = cfg.Default()

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
	Use:   "hc",
	Short: "hc is a tool to manage git hooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context(), args...)
	},
}

func runRoot(ctx context.Context, args ...string) error {
	cwd, err := config.Cwd()
	if err != nil {
		return err
	}

	cfg := filepath.Clean(filepath.Join(cwd, config.File))

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
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return err
		}
	}

	return nil
}
