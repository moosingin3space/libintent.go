package intent

import (
	"net/url"
	"regexp"
)

type Platform interface {
	Register(protocol string,
		path regexp.Regexp,
		handler chan<- Intent) (IntentReceiver, error)

	Unregister(recv IntentReceiver) error

	Send(url url.URL,
		metadata map[string]interface{},
		ack chan<- Intent) error
}

type IntentReceiver struct {
	Platform Platform
	protocol string
	path     regexp.Regexp
	done     <-chan bool
}

func (recv IntentReceiver) Unregister() {
	recv.Platform.Unregister(recv)
}
