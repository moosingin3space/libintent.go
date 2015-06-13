//+build linux, darwin

package registry

import (
	"net"
)

const socketPath = "/please/change/me"

type UnixRegistry struct {
	conn net.Conn
	quit <-chan bool
}

func RunRegistry() (*Registry, error) {
	socket, err := net.ListenUnixgram("unixgram", &net.UnixAddr{socketPath, "unixgram"})
	if err != nil {
		return nil, err
	}

	r := &UnixRegistry{
		conn: socket,
		quit: make(<-chan bool),
	}

	go server(socket, r.quit)

	return r, nil
}

func (r UnixRegistry) Terminate() {
	quit <- true
}

func server(conn net.Conn, quit <-chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			// TODO do work
		}
	}
}
