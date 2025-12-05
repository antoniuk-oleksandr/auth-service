package commoncaarlos0_env

import (
	commonenvparser "github.com/antoniuk-oleksandr/auth-service/common/env_parser"

	"github.com/caarlos0/env/v6"
)

type envParser struct{}

func New() commonenvparser.EnvParser {
	return &envParser{}
}

func (e *envParser) Parse(v any) error {
	return env.Parse(v)
}
