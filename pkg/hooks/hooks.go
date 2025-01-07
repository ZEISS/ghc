package hooks

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

var hookTemplate = "#!/usr/bin/env -S sh -c 'hc -c %s %s'"

// Path returns the path of the hook
func Path(ctx context.Context) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// check if the current directory is a git repository
	_, err = git.PlainOpen(filepath.Join(cwd, ".git"))
	if err != nil {
		return "", err
	}

	return filepath.Join(cwd, ".git", "hooks"), nil
}

// Install installs the hook
func Install(name, path, cfg string) error {
	tpl := fmt.Sprintf(hookTemplate, cfg, name)

	// nolint:gosec
	err := os.WriteFile(filepath.Join(path, name), []byte(tpl), 0o755)
	if err != nil {
		return err
	}

	return nil
}
