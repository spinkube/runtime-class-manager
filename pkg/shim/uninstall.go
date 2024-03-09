package shim

import (
	"errors"
	"log/slog"
	"os"

	"github.com/spinkube/runtime-class-manager/pkg/state"
)

func (c *Config) Uninstall(shimName string) (string, error) {
	st, err := state.Get(c.hostFs, c.kwasmPath)
	if err != nil {
		return "", err
	}
	s, ok := st.Shims[shimName]
	if !ok {
		slog.Warn("shim not installed", "shim", shimName)
		return "", nil
	}
	filePath := s.Path

	err = c.hostFs.Remove(filePath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return "", err
		}
		slog.Warn("shim binary did not exist, nothing to delete")
	}

	st.RemoveShim(shimName)
	if err := st.Write(); err != nil {
		return "", err
	}
	return filePath, err
}
