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

package containerd

import (
	"fmt"
	"log/slog"
	"syscall"

	"github.com/mitchellh/go-ps"
)

func RestartRuntime() error {
	pid, err := getPid()
	if err != nil {
		return err
	}
	slog.Info("found containerd pid", "pid", pid)

	err = syscall.Kill(pid, syscall.SIGHUP)
	if err != nil {
		return fmt.Errorf("failed to send SIGHUP to containerd: %+v", err)
	}
	return nil
}

func getPid() (int, error) {
	processList, err := ps.Processes()
	if err != nil {
		slog.Info("ps.Processes() Failed, are you using windows?")
		return -1, fmt.Errorf("could not get processes: %+v", err)
	}

	containerdProcessList := []ps.Process{}

	for x := range processList {
		process := processList[x]
		if process.Executable() == "containerd" {
			containerdProcessList = append(containerdProcessList, process)
		}
	}

	if len(containerdProcessList) == 1 {
		return containerdProcessList[0].Pid(), nil
	} else if len(containerdProcessList) == 0 {
		return -1, fmt.Errorf("no containerd process found")
	}

	panic("multiple containerd processes found")
}
