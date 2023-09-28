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
