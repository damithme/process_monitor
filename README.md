# Process Monitor

A mini `htop`-style CLI process monitor for macOS/Linux, built with Go stdlib only.

## Features

- List all running processes with PID, user, name, CPU%, and memory%
- Filter by CPU or memory threshold
- Kill a process by name or PID
- Auto-refresh display on a set interval
- Colour-coded output — red for high CPU, yellow for elevated CPU

## Usage

```bash
# List all processes
go run .

# Filter by CPU or memory
go run . --filter-cpu 5.0
go run . --filter-mem 1.0
go run . --filter-cpu 2.0 --filter-mem 0.5

# Kill a process
go run . --kill nginx
go run . --kill 1234

# Auto-refresh every N seconds (Ctrl+C to stop)
go run . --interval 3
go run . --interval 2 --filter-cpu 1.0
```

## Flags

| Flag | Type | Default | Description |
|---|---|---|---|
| `--filter-cpu` | float | `0.0` | Show only processes above this CPU% |
| `--filter-mem` | float | `0.0` | Show only processes above this MEM% |
| `--kill` | string | `""` | Kill process by name or PID |
| `--interval` | int | `0` | Auto-refresh every N seconds (0 = run once) |

## File Structure

```
process_monitor/
├── main.go       # Entry point, flag parsing, main loop
├── process.go    # Process struct, fetch + filter logic
├── display.go    # Colour table rendering
├── killer.go     # Kill by name or PID
└── go.mod
```

## Go Concepts Covered

| Concept | Where |
|---|---|
| Structs | `Process` type in `process.go` |
| `fmt.Printf` | Formatted table in `display.go` |
| `flag` package | CLI flags in `main.go` |
| `os/exec` | `ps aux` and `kill` commands |
| Pointers | `*Process` usage in killer functions |
| Goroutines | Background signal listener in `main.go` |
| Channels | `done chan struct{}` for clean shutdown |
| `time.Ticker` | Periodic refresh loop in `main.go` |

## Requirements

- Go 1.18+
- macOS or Linux (`ps aux` and `kill` must be available)

## Build

```bash
go build -o process-monitor .
./process-monitor --filter-cpu 2.0
```