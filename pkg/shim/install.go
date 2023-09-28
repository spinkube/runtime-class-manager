package shim

import (
	"io"
	"os"
	"path"
)

func Install(hostPath string, shimPath string, binaryDir string) (string, error) {
	srcFile, err := os.OpenFile(shimPath, os.O_RDONLY, 0000)
	if err != nil {
		return "", err
	}
	dstFilePath := path.Join(binaryDir, path.Base(shimPath))
	dstFilePathHost := path.Join(hostPath, dstFilePath)

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
