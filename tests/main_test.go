package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestMainUnsupportedDriver runs the main package as a subprocess with an
// unsupported antivirus driver to verify the process exits with an error.
func TestMainUnsupportedDriver(t *testing.T) {
	// Build path to the main.go
	mainPath := filepath.Join("cmd", "server", "main.go")

	cmd := exec.Command("go", "run", mainPath)
	// Set environment to use an unsupported driver so NewAntivirus returns error
	cmd.Env = append(os.Environ(), "ANTIVIRUS_DRIVER=unsupported_driver", "PORT=0")

	// Run and expect a non-zero exit (main should log.Fatalf)
	err := cmd.Run()
	if err == nil {
		t.Fatalf("expected process to exit with error when using unsupported driver")
	}
}
