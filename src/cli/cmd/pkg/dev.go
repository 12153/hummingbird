// Package pkg
package pkg 

import (
	"fmt"
	"os"
	"os/exec"
)

func Dev(projectDir string) error {
	fmt.Println("🚀 Starting dev environment...")

	cmds := []*exec.Cmd{
		exec.Command("templ", "generate"),
		exec.Command("air"),
		exec.Command("npm", "run", "dev"),
	}

	for _, cmd := range cmds {
		cmd.Dir = projectDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err := cmd.Start()
		if err != nil {
			return fmt.Errorf("failed to start process: %w", err)
		}
	}

	// Wait for all to finish
	for _, cmd := range cmds {
		cmd.Wait()
	}

	return nil
}
