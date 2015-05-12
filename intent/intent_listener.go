package intent

import (
	"encoding/binary"
	"github.com/tinylib/msgp/msgp"
	"io"
)

func intentListenerProc(conn io.ReadWriter,
	app Application,
	validator func(Intent) bool,
	handler chan<- Intent,
	quit <-chan bool) {

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

			// decode buffer using messagepack
			var data interface{}
			_, err = data.(msgp.Unmarshaler).UnmarshalMsg(buf)
			if err == nil {
				// nothing went wrong
				intent := data.(Intent)
				// validate the intent
				if validator(intent) {
					// tell it to handle the intent
					handler <- intent
				}
			}
		}
	}
}
