package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	PID    int
	Name   string
	CPU    float64
	Memory float64
	User   string
}

func FetchProcess() ([]Process, error) {
	cmd := exec.Command("ps", "aux")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch process list: %w", err)
	}
	lines := strings.Split(string(out), "\n")

	var processes []Process

	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) < 11 {
			continue
		}
		pid, _ := strconv.Atoi(fields[1])
		cpu, _ := strconv.ParseFloat(fields[2], 64)
		mem, _ := strconv.ParseFloat(fields[3], 64)
		name := fields[10]

		p := Process{
			PID:    pid,
			Name:   name,
			CPU:    cpu,
			Memory: mem,
			User:   fields[0],
		}
		processes = append(processes, p)
	}
	return processes, nil
}

func FilterProcesses(processes []Process, minCPU float64, minMem float64) []Process {
	var result []Process
	for _, p := range processes {
		if p.CPU <= minCPU && p.Memory <= minMem {
			result = append(result, p)
		}
	}
	return result
}
