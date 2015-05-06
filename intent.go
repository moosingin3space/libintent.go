package intent

func Register(protocol string,
	app Application,
	validator func(Intent) bool,
	handler <-chan Intent) (recv IntentReceiver, err error) {

	err := makeIntentDirectories()
	if err != nil {
		return
	}

	protocolSocket, err := makeProtocolSocket(protocol)
	if err != nil {
		return
	}

	// TODO kick off a goroutine to monitor this
}
