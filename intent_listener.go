package intent

import (
	"encoding/binary"
	"net"
)

func intentListenerProc(conn net.Conn,
	protocol string,
	app Application,
	validator func(Intent) bool,
	handler chan<- Intent,
	quit <-chan bool) {

	defer removeProtocolSocket(protocol)
	for {
		select {
		case <-quit:
			// quits when something is received
			return
		default:
			// The actual monitoring process
			var sizeBuf [4]byte
			n, err := conn.Read(sizeBuf[:])
			if err != nil {
				// fail hard (probably a better way)
				panic(err)
			}

			// Convert the buffer into a number
			size64, n := binary.Varint(sizeBuf[:])
			if n <= 0 {
				// this should never happen by protocol spec
				panic(err)
			}
			size := int(size64)

			// Now make a buffer to read in `size` bytes
			buf := make([]byte, size)
			n, err = conn.Read(buf[:])
			if err != nil || n != size {
				// something went wrong
				panic(err)
			}

			// TODO decode buffer using msgpack, validate the intent, and if passes, handle it
		}
	}
}
