//+build linux, darwin

package registry

import (
	"net"
)

type UnixRegistry struct {
	quit <-chan bool
}

func RunRegistry() (*Registry, error) {
	r := &UnixRegistry{
		quit: make(<-chan bool),
	}

	// TODO set up server

	go server(r.quit)

	return r, nil
}

func (r UnixRegistry) Terminate() {
	quit <- true
}

func server(quit <-chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			// TODO do work
		}
	}
}
