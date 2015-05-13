//go:generate msgp -o types_gen.go -io=false -tests=false
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
	protocol string
	done     chan bool
}
