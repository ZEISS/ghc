package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/zeiss/ghc/pkg/spec"

	"github.com/spf13/cobra"
)

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runValidate(cmd.Context())
	},
}

func runValidate(_ context.Context) error {
	path := filepath.Clean(config.File)

	s, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var spec spec.Spec
	if err := spec.UnmarshalYAML(s); err != nil {
		return err
	}

	return nil
}
