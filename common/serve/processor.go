package serve

type Processor interface {
	Init() error
	Driver() (string, interface{})
}
