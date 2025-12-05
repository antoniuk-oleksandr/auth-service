package hasher

type Hasher interface {
	Hash(value string) (string, error)
	Compare(value, hash string) error
}
