package auth

import (
	"crypto/rand"
	"encoding/hex"
	ports "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"
)

type keyGenerator struct {
}

func NewKeyGenerator() ports.KeyGenerator {
	return &keyGenerator{}
}

func (k *keyGenerator) Generate() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(key), nil
}
