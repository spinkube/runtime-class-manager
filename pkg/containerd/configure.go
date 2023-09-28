/*
   Copyright The KWasm Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

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
