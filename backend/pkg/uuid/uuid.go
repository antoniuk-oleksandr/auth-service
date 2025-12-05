package uuidgeneator

import (
	"github.com/google/uuid"
)

type UUIDGenerator interface {
	Generate() string
}

type uuidHelper struct{}

func New() UUIDGenerator {
	return &uuidHelper{}
}

func (u *uuidHelper) Generate() string {
	return uuid.NewString()
}
