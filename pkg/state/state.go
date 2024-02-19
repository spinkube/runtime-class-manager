package state

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"path"

	"github.com/kwasm/kwasm-node-installer/pkg/config"
)

type state struct {
	Shims  map[string]*Shim `json:"shims"`
	config *config.Config
}

func Get(config *config.Config) (*state, error) {
	out := state{
		Shims:  make(map[string]*Shim),
		config: config,
	}
	content, err := os.ReadFile(filePath(config))
	if err == nil {
		err := json.Unmarshal(content, &out)
		return &out, err
	}
	if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return &out, nil
}

func (l *state) ShimChanged(shimName string, sha256 []byte, path string) bool {
	shim, ok := l.Shims[shimName]
	if !ok {
		return true
	}

	return !bytes.Equal(shim.Sha256, sha256) || shim.Path != path
}

func (l *state) UpdateShim(shimName string, shim Shim) {
	l.Shims[shimName] = &shim
}

func (l *state) RemoveShim(shimName string) {
	delete(l.Shims, shimName)
}

func (l *state) Write() error {
	out, err := json.MarshalIndent(l, "", " ")
	if err != nil {
		return err
	}

	slog.Info("writing lock file", "content", string(out))

	return os.WriteFile(filePath(l.config), out, 0644)
}

func filePath(config *config.Config) string {
	return config.PathWithHost(path.Join(config.Kwasm.Path, "kwasm-lock.json"))
}
