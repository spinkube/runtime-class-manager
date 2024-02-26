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
	"log"
	"os"
	"path"
	"strings"

	"github.com/kwasm/kwasm-node-installer/pkg/config"
	"github.com/kwasm/kwasm-node-installer/pkg/shim"
	"github.com/spf13/afero"
)

type Config struct {
	config *config.Config
	fs     afero.Fs
}

func NewConfig(globalConfig *config.Config, fs afero.Fs) *Config {
	return &Config{
		config: globalConfig,
		fs:     fs,
	}
}

func (c *Config) AddRuntime(shimPath string) (string, error) {
	runtimeName := shim.RuntimeName(path.Base(shimPath))

	cfg := generateConfig(shimPath, runtimeName)

	configPath := configDirectory(c.config)
	configHostPath := c.config.PathWithHost(configPath)

	// Containerd config file needs to exist, otherwise return the error
	data, err := afero.ReadFile(c.fs, configHostPath)
	if err != nil {
		return configPath, err
	}

	// Fail if config.toml already contains the runtimeName
	// Prevents corrupt config but could lead to unexpcted fails for the user.
	// Maybe skipping existing config?
	if strings.Contains(string(data), runtimeName) {
		//return configPath, fmt.Errorf("config file %s already contains runtime config for '%s'", configPath, runtimeName)
		log.Printf("runtime '%s' already exists, skipping", runtimeName)
		return configPath, nil
	}

	// Open file in append mode
	file, err := c.fs.OpenFile(configHostPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return configPath, err
	}
	defer file.Close()

	// Append config
	_, err = file.WriteString(cfg)
	if err != nil {
		return configPath, err
	}

	return configPath, nil
}

func (c *Config) RemoveRuntime(shimPath string) (string, error) {
	runtimeName := shim.RuntimeName(path.Base(shimPath))

	cfg := generateConfig(shimPath, runtimeName)

	configPath := configDirectory(c.config)
	configHostPath := c.config.PathWithHost(configPath)

	// Containerd config file needs to exist, otherwise return the error
	data, err := afero.ReadFile(c.fs, configHostPath)
	if err != nil {
		return configPath, err
	}

	// Fail if config.toml does not contain the runtimeName
	if !strings.Contains(string(data), runtimeName) {
		return configPath, fmt.Errorf("config file %s does not contain a runtime config for '%s'", configPath, runtimeName)
	}

	// Convert the file data to a string and replace the target string with an empty string.
	modifiedData := strings.Replace(string(data), cfg, "", -1)

	// Write the modified data back to the file.
	err = afero.WriteFile(c.fs, configHostPath, []byte(modifiedData), 0644)
	if err != nil {
		log.Fatal(err)
	}

	return configPath, nil
}

func configDirectory(config *config.Config) string {
	return config.Runtime.ConfigPath
}

func generateConfig(shimPath string, runtimeName string) string {
	return fmt.Sprintf(`
# KWASM runtime config for %s
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.%s]
runtime_type = "%s"
`, runtimeName, runtimeName, shimPath)
}
