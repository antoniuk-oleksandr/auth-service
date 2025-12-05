package http

type HTTPRouter interface {
	Start(address string) error
	Group(path string) RouterGroup
	Close() error
}
