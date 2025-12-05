package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost int) Hasher {
	return &bcryptHasher{
		cost: cost,
	}
}

func (h *bcryptHasher) Compare(value string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
}

func (h *bcryptHasher) Hash(value string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(value), h.cost)
	if err != nil {
		return "", err
	}
	
	return string(result), nil
}
