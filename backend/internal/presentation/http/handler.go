package http

type Handler func(ctx HTTPContext) error
