package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/zeiss/ghc/pkg/hooks"
	"github.com/zeiss/ghc/pkg/spec"

	"github.com/spf13/cobra"
)

var UninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: `Uninstall the git hooks`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runUninstall(cmd.Context())
	},
}

func runUninstall(ctx context.Context) error {
	cfg := filepath.Clean(config.File)

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
		if err := hooks.Uninstall(name, path); err != nil {
			return err
		}
	}

	return nil
}
