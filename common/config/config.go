package commonconfig

import (
	envparser "github.com/antoniuk-oleksandr/auth-service/common/env_parser"
	val "github.com/antoniuk-oleksandr/auth-service/common/validator"
)

func LoadAppConfig(v val.Validator, parser envparser.EnvParser, cfg any) error {
	if err := parser.Parse(cfg); err != nil {
		return err
	}

	if err := v.Struct(cfg); err != nil {
		return err
	}

	return nil
}
