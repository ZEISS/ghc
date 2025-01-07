package hooks

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/zeiss/pkg/conv"
)

var hookTemplate = "#!/usr/bin/env -S sh -c 'ghc -c %s -r %s'"

// Path returns the path of the hook
func Path(ctx context.Context) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	cmd := exec.CommandContext(ctx, "git", "config", "core.hooksPath")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return filepath.Clean(filepath.Join(cwd, strings.TrimSuffix(conv.String(out), "\n"))), nil
}

// Install installs the hook
func Install(name, path, cfg string) error {
	tpl := fmt.Sprintf(hookTemplate, cfg, name)

	// nolint:gosec
	err := os.WriteFile(filepath.Clean(filepath.Join(path, name)), []byte(tpl), 0o755)
	if err != nil {
		return err
	}

	return nil
}
