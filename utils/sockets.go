package utils

import (
	socketio "github.com/googollee/go-socket.io"
)

// SocketServer definition
var SocketServer *socketio.Server

// InitSocket handler for use by app
func InitSocket() error {
	var err error
	SocketServer, err = socketio.NewServer(nil)

	if err != nil {
		return err
	}

	return nil
}
