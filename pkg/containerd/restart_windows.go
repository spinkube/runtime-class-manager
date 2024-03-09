//go:build windows
// +build windows

package containerd

import "errors"

type ContainerdRestarter struct{}

func (r ContainerdRestarter) Restart() error {
	return errors.New("restarting containerd not implemented")
}
