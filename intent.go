package intent

func Register(protocol string,
	app Application,
	validator func(Intent) bool,
	handler <-chan Intent) (recv IntentReceiver, err error) {

	done := make(chan bool)
	err := makeIntentDirectories()
	if err != nil {
		return
	}

	protocolSocket, err := makeProtocolSocket(protocol)
	if err != nil {
		return
	}

	go intentListenerProc(protocolSocket, protocol, app, validator, handler, done)
	return &IntentReceiver{done: done}
}

func Unregister(recv IntentReceiver) {
	recv.done <- true
}
