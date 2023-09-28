package containerd

import (
	"fmt"
	"os"
	"path"

	"github.com/kwasm/kwasm-node-installer/pkg/config"
	"github.com/kwasm/kwasm-node-installer/pkg/shim"
)

func WriteConfig(config *config.Config, shimPath string) (string, error) {
	runtimeName := shim.RuntimeName(path.Base(shimPath))

	cfg := generateConfig(shimPath, runtimeName)

	configPath := path.Join(configDirectory(config), fmt.Sprintf("%s.%s", runtimeName, "toml"))
	configHostPath := config.PathWithHost(configPath)

	err := os.MkdirAll(path.Dir(configHostPath), 0755)
	if err != nil {
		return configPath, err
	}

	return configPath, os.WriteFile(configHostPath, []byte(cfg), 0644)
}

func configDirectory(config *config.Config) string {
	return path.Join(path.Dir(config.Runtime.ConfigPath), "conf.d")
}

func generateConfig(shimPath string, runtimeName string) string {
	return fmt.Sprintf(`[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.%s]
runtime_type = "%s"`, runtimeName, shimPath)
}
