package shim

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/spinkube/runtime-class-manager/internal/state"
)

func (c *Config) Uninstall(shimName string) (string, error) {
	st, err := state.Get(c.hostFs, c.kwasmPath)
	if err != nil {
		return "", err
	}
	s, ok := st.Shims[shimName]
	if !ok {
		slog.Error("shim not installed", "shim", shimName)
		return "", fmt.Errorf("shim %s not installed", shimName)
	}
	filePath := s.Path

	err = c.hostFs.Remove(filePath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("shim binary did not exist, nothing to delete")
		}
	}
	st.RemoveShim(shimName)
	if err = st.Write(); err != nil {
		return "", err
	}
	return filePath, err
}
