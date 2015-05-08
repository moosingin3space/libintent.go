// go:generate msgp
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
	done chan bool
}
