package ui

import (
	"fmt"

	"github.com/acousticclown/ports/ports"
)

func PrintTable(portsList []ports.PortInfo) {
	Title.Println("\n🚀 Running Ports\n")

	Header.Printf("%-12s %-8s %-8s %-6s\n", "PROCESS", "PID", "PROTO", "PORT")
	Muted.Println("---------------------------------------------")

	for _, p := range portsList {
		fmt.Printf(
			"%-12s %-8s %-8s %-6s\n",
			p.Process, p.PID, p.Protocol, p.Port,
		)
	}
	fmt.Println()
}
