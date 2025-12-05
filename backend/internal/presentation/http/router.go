package http

type HTTPRouter interface {
	Start(address string) error
	Use(middleware any)
	Group(path string) RouterGroup
	Close() error
}
