package main

import (
	"fmt"
	"os"

	"github.com/acousticclown/ports/ports"
	"github.com/acousticclown/ports/ui"
)

func printHelp() {
	fmt.Println(`
		Ports 🔌
		A fast and minimal CLI to list, check, and kill running ports.

		Usage:
		ports                 List all running ports
		ports check <port>    Check if a port is in use
		ports kill <port>     Kill process running on a port
		ports version         Show version information
		ports help            Show this help

		Flags:
		-h, --help            Show help
		-v, --version         Show version

		Examples:
		ports
		ports check 8080
		ports kill 3000
		ports version
	`)
}

func printVersion() {
	fmt.Printf(
		"ports %s\ncommit: %s\nbuilt:  %s\n",
		Version, Commit, Date,
	)
}

func main() {
	args := os.Args

	// No args → default behavior
	if len(args) == 1 {
		list, err := ports.ListPorts()
		if err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}
		ui.PrintTable(list)
		return
	}

	switch args[1] {

	case "help", "--help", "-h":
		printHelp()
		return

	case "version", "--version", "-v":
		printVersion()
		return

	case "check":
		if len(args) != 3 {
			ui.PrintError("Usage: ports check <port>")
			os.Exit(1)
		}
		running, info, err := ports.IsPortRunning(args[2])
		if err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}
		if running {
			ui.PrintSuccess(
				fmt.Sprintf(
					"Port %s is running (Process: %s, PID: %s)",
					args[2], info.Process, info.PID,
				),
			)
		} else {
			ui.PrintError("Port " + args[2] + " is NOT running")
		}
		return

	case "kill":
		if len(args) != 3 {
			ui.PrintError("Usage: ports kill <port>")
			os.Exit(1)
		}
		err := ports.KillPort(args[2])
		if err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}
		ui.PrintSuccess("Port " + args[2] + " killed successfully")
		return

	default:
		ui.PrintError("Unknown command: " + args[1])
		printHelp()
		os.Exit(1)
	}
}
