package http

type Handler func(c HTTPContext) error
