package types

type Protocol interface {
	RegisterHandler(command string, handlers ...HandlerFunc)
	Start(port string) error
}
