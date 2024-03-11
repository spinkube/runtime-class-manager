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

package main_test

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/spinkube/runtime-class-manager/cmd/node-installer"
	"github.com/spinkube/runtime-class-manager/tests"
	"github.com/stretchr/testify/require"
)

type nullRestarter struct{}

func (n nullRestarter) Restart() error {
	return nil
}

func Test_RunInstall(t *testing.T) {
	type args struct {
		config cmd.Config
		rootFs afero.Fs
		hostFs afero.Fs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"new shim",
			args{
				cmd.Config{
					struct {
						Name       string
						ConfigPath string
					}{"containerd", "/etc/containerd/config.toml"},
					struct {
						Path      string
						AssetPath string
					}{"/opt/kwasm", "/assets"},
					struct{ RootPath string }{"/containerd/missing-containerd-shim-config"},
				},
				tests.FixtureFs("../testdata"),
				tests.FixtureFs("../testdata/containerd/missing-containerd-shim-config"),
			},
			false,
		},
		{
			"existing shim",
			args{
				cmd.Config{
					struct {
						Name       string
						ConfigPath string
					}{"containerd", "/etc/containerd/config.toml"},
					struct {
						Path      string
						AssetPath string
					}{"/opt/kwasm", "/assets"},
					struct{ RootPath string }{"/containerd/existing-containerd-shim-config"},
				},
				tests.FixtureFs("../testdata"),
				tests.FixtureFs("../testdata/containerd/existing-containerd-shim-config"),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cmd.RunInstall(tt.args.config, tt.args.rootFs, tt.args.hostFs, nullRestarter{})
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
