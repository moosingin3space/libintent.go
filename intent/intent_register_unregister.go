package intent

func Register(platform Platform,
	protocol string,
	app Application,
	validator func(Intent) bool,
	handler chan<- Intent) (recv IntentReceiver, err error) {

	done := make(chan bool)
	protocolSocket, err := platform.NewSocket(platform.Protocol(protocol))
	if err != nil {
		return
	}

	go intentListenerProc(protocolSocket, app, validator, handler, done)
	recv = IntentReceiver{done: done, protocol: protocol}
	return
}

func Unregister(platform Platform, recv IntentReceiver) {
	recv.done <- true
	platform.CleanupSocket(platform.Protocol(recv.protocol))
}
