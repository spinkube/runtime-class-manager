package shim

import "strings"

func RuntimeName(bin string) string {
	return strings.TrimPrefix(bin, "containerd-shim-")
}
