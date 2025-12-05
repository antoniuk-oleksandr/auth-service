package commonconfig

import (
	envparser "auth-common/env_parser"
	val "auth-common/validator"
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
