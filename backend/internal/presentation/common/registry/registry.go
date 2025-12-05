package registry

type Routes interface {
	Register()
}

type Registry interface {
	Add(routes Routes) Registry
	Register()
}

type registry struct {
	routeGroups []Routes
}

func New() Registry {
	return &registry{}
}

func (r *registry) Register() {
	for _, rg := range r.routeGroups {
		rg.Register()
	}
}

func (r *registry) Add(group Routes) Registry {
	r.routeGroups = append(r.routeGroups, group)
	return r
}
