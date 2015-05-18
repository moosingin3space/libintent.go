package intent

type Verb int

const (
	READ Verb = iota
	WRITE
	DELETE
	CAPS
)

type Intent struct {
	Protocol string
	Path     string
	Verb     Verb
	Metadata map[string]interface{}
	Payload  []byte
}
