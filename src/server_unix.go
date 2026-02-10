//go:build !windows

package richpresence

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func createListener(index int) (net.Listener, error) {
	socketPath := getSocketPath(index)
	if socketPath == "" {
		return nil, fmt.Errorf("could not determine socket path")
	}

	// remove the old one (discord may be open(?))
	os.Remove(socketPath)

	// ensure exists
	dir := filepath.Dir(socketPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	return net.Listen("unix", socketPath)
}

func getSocketPath(index int) string {
	socketName := fmt.Sprintf("discord-ipc-%d", index)

	// XDG_RUNTIME_DIR (preferred on Linux)
	if xdgRuntime := os.Getenv("XDG_RUNTIME_DIR"); xdgRuntime != "" {
		return filepath.Join(xdgRuntime, socketName)
	}

	// TMPDIR (common on macOS)
	if tmpDir := os.Getenv("TMPDIR"); tmpDir != "" {
		return filepath.Join(tmpDir, socketName)
	}

	// /run/user/{uid} (fallback for Linux)
	if uid := os.Getuid(); uid >= 0 {
		runUserPath := fmt.Sprintf("/run/user/%d", uid)
		if _, err := os.Stat(runUserPath); err == nil {
			return filepath.Join(runUserPath, socketName)
		}
	}

	// /tmp (last resort)
	return filepath.Join("/tmp", socketName)
}
