package shim

import (
	"fmt"
	"os"

	"github.com/spinkube/runtime-class-manager/pkg/config"
	"github.com/spinkube/runtime-class-manager/pkg/state"
)

func Uninstall(config *config.Config, shimName string) (string, error) {
	st, err := state.Get(config)
	if err != nil {
		return "", fmt.Errorf("error getting state: %w", err)
	}
	s := st.Shims[shimName]
	if s == nil {
		return "", fmt.Errorf("shim '%s' not installed", shimName)
	}
	filePath := s.Path
	filePathHost := config.PathWithHost(filePath)

	err = os.Remove(filePathHost)
	if err != nil {
		return "", fmt.Errorf("error removing shim: %w", err)
	}

	st.RemoveShim(shimName)
	if err := st.Write(); err != nil {
		return "", fmt.Errorf("error writing state: %w", err)
	}
	return filePath, fmt.Errorf("error uninstalling shim: %w", err)
}
