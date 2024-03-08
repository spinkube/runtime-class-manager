package state

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/spinkube/runtime-class-manager/pkg/config"
)

type State struct {
	Shims  map[string]*Shim `json:"shims"`
	config *config.Config
}

func Get(config *config.Config) (*State, error) {
	out := State{
		Shims:  make(map[string]*Shim),
		config: config,
	}
	content, err := os.ReadFile(filePath(config))
	if err == nil {
		err := json.Unmarshal(content, &out)
		return &out, fmt.Errorf("error reading file: %w", err)
	}
	if !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return &out, nil
}

func (l *State) ShimChanged(shimName string, sha256 []byte, path string) bool {
	shim, ok := l.Shims[shimName]
	if !ok {
		return true
	}

	return !bytes.Equal(shim.Sha256, sha256) || shim.Path != path
}

func (l *State) UpdateShim(shimName string, shim Shim) {
	l.Shims[shimName] = &shim
}

func (l *State) RemoveShim(shimName string) {
	delete(l.Shims, shimName)
}

func (l *State) Write() error {
	out, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		return fmt.Errorf("error marshalling state: %w", err)
	}

	slog.Info("writing lock file", "content", string(out))

	return os.WriteFile(filePath(l.config), out, 0o644)
}

func filePath(config *config.Config) string {
	return config.PathWithHost(path.Join(config.Kwasm.Path, "kwasm-lock.json"))
}
