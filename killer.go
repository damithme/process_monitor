package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func KillByPID(pid int) error {
	cmd := exec.Command("kill", strconv.Itoa(pid))
	return cmd.Run()
}

func KillByName(name string, process []Process) error {
	for _, process := range process {
		if process.Name == name {
			return KillByPID(process.PID)
		}
	}
	return fmt.Errorf("no process found with name: %s", name)
}

func KillTarget(target string, process []Process) error {
	pid, err := strconv.Atoi(target)
	if err != nil {
		return KillByPID(pid)
	}
	return KillByName(target, process)
}
