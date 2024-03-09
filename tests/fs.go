package tests

import (
	"github.com/spf13/afero"
)

func FixtureFs(fixturePath string) afero.Fs {
	baseFs := afero.NewBasePathFs(afero.NewOsFs(), fixturePath)
	fs := afero.NewCopyOnWriteFs(baseFs, afero.NewMemMapFs())
	return fs
}
