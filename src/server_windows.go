//go:build windows

package richpresence

import (
	"fmt"
	"net"

	"gopkg.in/natefinch/npipe.v2"
)

func createListener(index int) (net.Listener, error) {
	pipePath := fmt.Sprintf(`\\.\pipe\discord-ipc-%d`, index)
	return npipe.Listen(pipePath)
}
