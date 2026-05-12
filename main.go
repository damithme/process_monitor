package main

import (
	"flag"
	"fmt"
)

func PrintHeader() {
	fmt.Printf("%-9s %-16s %-29s %7s %7s\n",
		"PID", "USER", "NAME", "CPU%", "MEM%")
	fmt.Printf("%-9s %-16s %-29s %7s %7s\n",
		"---------", "----------------", "-----------------------------", "-------", "-------")
}

func PrintProcess(p Process) {
	name := p.Name
	if len(name) > 29 {
		name = name[:26] + "..."
	}

	fmt.Printf("%-9d %-16s %-29s %7.1f %7.1f\n",
		p.PID, p.User, name, p.CPU, p.Memory)
}

func PrintAll(processes []Process) {
	PrintHeader()
	for _, p := range processes {
		PrintProcess(p)
	}
}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	filterCPU := flag.Float64("filter-cpu", 0.0, "Show processes above this CPU%")
	filterMem := flag.Float64("filter-mem", 0.0, "Show processes above this MEM%")
	killTarget := flag.String("kill", "", "kill process by name or PID")
	interval := flag.Int("interval", 0, "auto-refresh every N seconds (0 = run once")

	flag.Parse()

	processes, err := FetchProcess()

	if *killTarget != "" {
		err := KillTarget(*killTarget, processes)
		if err != nil {
			fmt.Println("Kill failed:", err)
		} else {
			fmt.Println("Killed", *killTarget)
		}
		return
	}
	_ = interval // not wired up yet

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	filtered := FilterProcesses(processes, *filterCPU, *filterMem)
	PrintAll(filtered)
}
