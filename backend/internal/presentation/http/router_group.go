package http

type RouterGroup interface {
	Get(path string, handler Handler) RouterGroup
	Post(path string, handler Handler) RouterGroup
	Put(path string, handler Handler) RouterGroup
	Delete(path string, handler Handler) RouterGroup
}
