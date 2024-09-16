package preset

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/spf13/afero"
	"github.com/spinkube/runtime-class-manager/internal/containerd"
)

type Settings struct {
	ConfigPath string
	Setup      func(Env) error
	Restarter  func(Env) containerd.Restarter
}

type Env struct {
	HostFs     afero.Fs
	ConfigPath string
}

var Default = Settings{
	ConfigPath: "/etc/containerd/config.toml",
	Setup:      func(_ Env) error { return nil },
	Restarter:  func(_ Env) containerd.Restarter { return containerd.NewRestarter() },
}

func (s Settings) WithConfigPath(path string) Settings {
	s.ConfigPath = path
	return s
}

func (s Settings) WithSetup(setup func(env Env) error) Settings {
	s.Setup = setup
	return s
}

var MicroK8s = Default.WithConfigPath("/var/snap/microk8s/current/args/containerd-template.toml")

var RKE2 = Default.WithConfigPath("/var/lib/rancher/rke2/agent/etc/containerd/config.toml.tmpl").
	WithSetup(func(env Env) error {
		_, err := env.HostFs.Stat(env.ConfigPath)
		if err == nil {
			return nil
		}

		if errors.Is(err, os.ErrNotExist) {
			// Copy base config into .tmpl version
			src, _ := strings.CutSuffix(env.ConfigPath, ".tmpl")
			in, err := env.HostFs.Open(src)
			if err != nil {
				return err
			}
			defer in.Close()
			out, err := env.HostFs.Create(env.ConfigPath)
			if err != nil {
				return err
			}
			defer func() {
				cerr := out.Close()
				if err == nil {
					err = cerr
				}
			}()
			if _, err = io.Copy(out, in); err != nil {
				return err
			}
			err = out.Sync()

			return nil
		}

		return err
	})

var K3s = RKE2.WithConfigPath("/var/lib/rancher/k3s/agent/etc/containerd/config.toml.tmpl")
