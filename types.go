package intent

type Intent struct {
	Url    string
	Params map[string][]byte
}

type Application struct {
	Name    string
	Version string
}

type IntentReceiver struct {
	incomingIntents chan Intent
	done            chan bool
}
