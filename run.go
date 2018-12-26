package main

import (
	"os/exec"
	"os"
)

// Run a shell command and send output to stdout
// and stderr.
func runCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}