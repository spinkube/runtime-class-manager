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
	"path"

	"github.com/spf13/cobra"

	"github.com/kwasm/kwasm-node-installer/pkg/containerd"
	"github.com/kwasm/kwasm-node-installer/pkg/shim"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall containerd shims",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("uninstall called", "config", config)
		shimName := config.Runtime.Name
		runtimeName := path.Join(config.Kwasm.Path, "bin", shimName)

		binPath, err := shim.Uninstall(&config, shimName)
		if err != nil {
			slog.Error("failed to uninstall shim", "shim", runtimeName, "error", err)
			return
		}

		configPath, err := containerd.RemoveRuntime(&config, binPath)
		if err != nil {
			slog.Error("failed to write containerd config", "shim", runtimeName, "path", configPath, "error", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
