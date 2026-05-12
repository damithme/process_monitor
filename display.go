package main

import "fmt"

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	cyan   = "\033[36m"
	yellow = "\033[33m"
	red    = "\033[31m"
	dim    = "\033[2m"
)

func cpuColor(cpu float64) string {
	switch {
	case cpu >= 10.0:
		return red
	case cpu >= 5.0:
		return yellow
	default:
		return reset
	}
}

func PrintHeader() {
	fmt.Printf("%s%-9s %-16s %-29s %7s %7s%s\n",
		bold+cyan, "PID", "USER", "NAME", "CPU%", "MEM%", reset)
	fmt.Printf("%s%-9s %-16s %-29s %7s %7s%s\n",
		dim, "---------", "----------------", "-----------------------------", "-------", "-------", reset)
}

func PrintProcess(p Process) {
	name := p.Name
	if len(name) > 29 {
		name = name[:26] + "..."
	}
	color := cpuColor(p.CPU)
	fmt.Printf("%s%-9d %-16s %-29s %7.1f %7.1f%s\n",
		color, p.PID, p.User, name, p.CPU, p.Memory, reset)
}

func PrintAll(processes []Process) {
	PrintHeader()
	for _, p := range processes {
		PrintProcess(p)
	}
	fmt.Printf("%s  %d processes%s\n", dim, len(processes), reset)
}
