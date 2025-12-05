package envparser

import (
	"github.com/caarlos0/env/v6"
)

type envParser struct{}

func New() EnvParser {
	return &envParser{}
}

func (e *envParser) Parse(v any) error {
	return env.Parse(v)
}
