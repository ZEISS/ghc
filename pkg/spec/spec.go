package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/zeiss/pkg/filex"
	"github.com/zeiss/pkg/mapx"
	"gopkg.in/yaml.v3"
)

// ErrHookNotExist is an error when the hook does not exist
var ErrHookNotExist = errors.New("hook does not exist")

// GitHooks is a list of git hooks
// See: https://git-scm.com/docs/githooks
var GitHooks = []string{
	"applypatch-msg",
	"commit-msg",
	"post-applypatch",
	"post-checkout",
	"post-commit",
	"post-merge",
	"post-rewrite",
	"pre-applypatch",
	"pre-auto-gc",
	"pre-commit",
	"pre-merge-commit",
	"pre-push",
	"pre-rebase",
	"prepare-commit-msg",
}

// Example returns an example of the Spec struct
func Example() *Spec {
	return &Spec{
		Version:     DefaultVersion,
		Name:        "example",
		Description: "This is an example of the specification",
		Stdout:      true,
		Stderr:      true,
		Hooks: Hooks{
			"pre-commit": []string{"go test ./..."},
		},
	}
}

const (
	// DefaultVersion is the default version of the specification
	DefaultVersion = 1
	// DefaultFilename is the default filename of the specification
	DefaultFilename = ".ghc.yaml"
)

var validate = validator.New()

// Spec is the specification of the repository
type Spec struct {
	// Version is the version of the specification
	Version int `yaml:"version" validate:"required"`
	// Name is a given name of the project or repository (optional)
	Name string `yaml:"name,omitempty"`
	// Description is a short description of the project or repository (optional)
	Description string `yaml:"description,omitempty"`
	// Hooks is a list of hooks to be executed
	Hooks Hooks `yaml:"hooks" validate:"required"`
	// Stdout is a list of hooks to be executed
	Stdout bool `yaml:"stdout,omitempty"`
	// Stderr is a list of hooks to be
	Stderr bool `yaml:"stderr,omitempty"`

	sync.Mutex `yaml:"-"`
}

// Hooks is a `git` hook
type Hooks map[string][]string

// Hook returns the hook with the given name
func (s *Spec) Hook(name string) ([]string, error) {
	s.Lock()
	defer s.Unlock()

	if !mapx.Exists(s.Hooks, name) {
		return nil, ErrHookNotExist
	}

	return s.Hooks[name], nil
}

// UnmarshalYAML unmarshals the YAML data into the Spec struct
func (s *Spec) UnmarshalYAML(data []byte) error {
	ss := struct {
		Version     int    `yaml:"version"`
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Hooks       Hooks  `yaml:"hooks"`
	}{}

	if err := yaml.Unmarshal(data, &ss); err != nil {
		return errors.WithStack(err)
	}

	s.Version = ss.Version
	s.Name = ss.Name
	s.Description = ss.Description
	s.Hooks = ss.Hooks

	err := validate.Struct(s)
	if err != nil {
		return err
	}

	return err
}

// Write writes the specification to the given file.
func Write(s *Spec, file string, force bool) error {
	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	ok, _ := filex.FileExists(filepath.Clean(file))
	if ok && !force {
		return fmt.Errorf("%s already exists, use --force to overwrite", file)
	}

	f, err := os.Create(filepath.Clean(file))
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
