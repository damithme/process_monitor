package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	filterCPU := flag.Float64("filter-cpu", 0.0, "Show processes above this CPU%")
	filterMem := flag.Float64("filter-mem", 0.0, "Show processes above this MEM%")
	killTarget := flag.String("kill", "", "kill process by name or PID")
	interval := flag.Int("interval", 0, "auto-refresh every N seconds (0 = run once)")

	flag.Parse()

	if *killTarget != "" {
		processes, err := FetchProcess()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = KillTarget(*killTarget, processes)
		if err != nil {
			fmt.Println("Kill failed:", err)
		} else {
			fmt.Println("Killed", *killTarget)
		}
		return
	}

	if *interval > 0 {
		ticker := time.NewTicker(time.Duration(*interval) * time.Second)
		defer ticker.Stop()

		done := make(chan struct{})

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sigs
			close(done)
		}()

		for {
			select {
			case <-ticker.C:
				processes, err := FetchProcess()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				filtered := FilterProcesses(processes, *filterCPU, *filterMem)
				fmt.Print("\033[H\033[2J")
				PrintAll(filtered)

			case <-done:
				fmt.Println("\nStopped.")
				return
			}
		}
	} else {
		processes, err := FetchProcess()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		filtered := FilterProcesses(processes, *filterCPU, *filterMem)
		PrintAll(filtered)
	}
}
