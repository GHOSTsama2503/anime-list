package common

import (
	"os/exec"
	"strings"
)

func Version() string {
	cmd := exec.Command("git", "describe", "--tags")

	out, err := cmd.Output()
	if err != nil {
		return "v0.0.1"
	}

	return strings.TrimSpace(string(out))
}
