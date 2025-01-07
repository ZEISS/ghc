package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/zeiss/ghc/pkg/spec"

	"github.com/spf13/cobra"
	"github.com/zeiss/pkg/filex"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new config",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit(cmd.Context())
	},
}

func runInit(_ context.Context) error {
	if err := spec.Write(spec.Example(), spec.DefaultFilename, config.Force); err != nil {
		return err
	}

	cwd, err := config.Cwd()
	if err != nil {
		return err
	}

	path := filepath.Clean(filepath.Join(cwd, spec.DefaultFilename))

	err = filex.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
