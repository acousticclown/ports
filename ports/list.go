package ports

import (
	"os/exec"
	"strings"
)

func ListPorts() ([]PortInfo, error) {
	cmd := exec.Command("lsof", "-i", "-P", "-n")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var results []PortInfo

	for i, line := range lines {
		if i == 0 || strings.TrimSpace(line) == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 9 {
			continue
		}

		// Example:
		// node  12345 user  ... TCP *:3000 (LISTEN)
		address := fields[len(fields)-2]
		if !strings.Contains(address, ":") {
			continue
		}

		port := strings.Split(address, ":")[1]

		results = append(results, PortInfo{
			Process: fields[0],
			PID:     fields[1],
			Protocol: fields[7],
			Port:    port,
		})
	}

	return results, nil
}
