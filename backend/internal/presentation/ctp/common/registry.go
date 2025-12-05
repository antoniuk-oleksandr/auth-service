package common

type CTPRoutes interface {
	Register()
}

type Registry interface {
	Add(item any)
}