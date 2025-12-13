package main

import (
	"fmt"
	"os"

	"github.com/acousticclown/ports/ports"
	"github.com/acousticclown/ports/ui"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		// ports
		list, err := ports.ListPorts()
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}
		ui.PrintTable(list)
		return
	}

	if len(args) == 3 && args[1] == "check" {
		port := args[2]
		running, info, err := ports.IsPortRunning(port)
		if err != nil {
			fmt.Println("❌ Error:", err)
			return
		}

		if running {
			fmt.Printf("✅ Port %s is running (Process: %s, PID: %s)\n",
				port, info.Process, info.PID)
		} else {
			fmt.Printf("❌ Port %s is NOT running\n", port)
		}
		return
	}

	if len(args) == 3 && args[1] == "kill" {
	port := args[2]
	err := ports.KillPort(port)
	if err != nil {
		ui.PrintError(err.Error())
		return
	}
	ui.PrintSuccess("Port " + port + " killed successfully")
	return
}


	fmt.Println("Usage:")
	fmt.Println("  ports")
	fmt.Println("  ports check <port>")
}
