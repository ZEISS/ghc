package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/zeiss/ghc/pkg/hooks"
	"github.com/zeiss/ghc/pkg/spec"

	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the git hooks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInstall(cmd.Context())
	},
}

func runInstall(ctx context.Context) error {
	cwd, err := config.Cwd()
	if err != nil {
		return err
	}

	cfg := config.File
	if !filepath.IsAbs(config.File) {
		cfg = filepath.Clean(filepath.Join(cwd, config.File))
	}

	s, err := os.ReadFile(cfg)
	if err != nil {
		return err
	}

	var spec spec.Spec
	if err := spec.UnmarshalYAML(s); err != nil {
		return err
	}

	path, err := hooks.Path(ctx)
	if err != nil {
		return err
	}

	for name := range spec.Hooks {
		if err := hooks.Install(name, path, cfg); err != nil {
			return err
		}
	}

	return nil
}
