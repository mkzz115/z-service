package serve

type Serve interface {
	Init() error
	Start() error
}
