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

	"github.com/spf13/cobra"

	"github.com/kwasm/kwasm-node-installer/pkg/containerd"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install containerd shims",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("install called", "config", config)

		err := containerd.RestartRuntime()
		if err != nil {
			slog.Error("failed to restart containerd", "error", err)
		}
	},
}

func init() {
	installCmd.Flags().StringVarP(&config.Kwasm.AssetPath, "asset-path", "a", "/assets", "Path to the binaries and libraries to install")
	installCmd.Flags().StringVarP(&config.Runtime.CRIPluginName, "cri-plugin-name", "p", "\"io.containerd.grpc.v1.cri\"", "Name of the cri plugin")
	rootCmd.AddCommand(installCmd)
}
