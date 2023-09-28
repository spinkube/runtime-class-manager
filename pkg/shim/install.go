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

package shim

import (
	"io"
	"os"
	"path"

	"github.com/kwasm/kwasm-node-installer/pkg/config"
)

func Install(config *config.Config, shimName string) (string, error) {
	shimPath := config.AssetPath(shimName)
	srcFile, err := os.OpenFile(shimPath, os.O_RDONLY, 0000)
	if err != nil {
		return "", err
	}
	dstFilePath := path.Join(config.Kwasm.Path, "bin", shimName)
	dstFilePathHost := config.PathWithHost(dstFilePath)

	err = os.MkdirAll(path.Dir(dstFilePathHost), 0755)
	if err != nil {
		return dstFilePath, err
	}

	dstFile, err := os.OpenFile(dstFilePathHost, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(dstFile, srcFile)

	return dstFilePath, err
}
