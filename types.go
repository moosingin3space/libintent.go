package intent

type Intent struct {
	Url    string
	Params map[int8]string
}

type Application struct {
	Name    string
	Version string
}
