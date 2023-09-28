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

package cmd

import (
	"log/slog"
	"os"

	"github.com/kwasm/kwasm-node-installer/pkg/containerd"
	"github.com/kwasm/kwasm-node-installer/pkg/shim"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install containerd shims",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := os.ReadDir(config.Kwasm.AssetPath)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		for _, file := range files {
			binPath, err := shim.Install(&config, file.Name())
			if err != nil {
				slog.Error(err.Error())
				return
			}
			slog.Info("shim installed", "shim", shim.RuntimeName(file.Name()), "path", binPath)
			configPath, err := containerd.WriteConfig(&config, binPath)
			if err != nil {
				slog.Error(err.Error())
				return
			}
			slog.Info("shim configured", "shim", shim.RuntimeName(file.Name()), "path", configPath)
		}

		err = containerd.RestartRuntime()
		if err != nil {
			slog.Error("failed to restart containerd", "error", err)
		}
	},
}

func init() {
	installCmd.Flags().StringVarP(&config.Kwasm.AssetPath, "asset-path", "a", "/assets", "Path to the binaries and libraries to install")
	rootCmd.AddCommand(installCmd)
}
