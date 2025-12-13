package ports

import (
	"fmt"
	"os/exec"
)

func KillPort(port string) error {
	// Get PID using lsof
	cmd := exec.Command("lsof", "-t", "-i", ":"+port)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("no process found on port %s", port)
	}

	pid := string(output)
	pid = pid[:len(pid)-1] // remove newline

	// Kill process
	killCmd := exec.Command("kill", "-9", pid)
	err = killCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to kill process on port %s", port)
	}

	return nil
}
