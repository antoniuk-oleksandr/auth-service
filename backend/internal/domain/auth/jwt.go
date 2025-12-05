package auth

type JWT struct {
	AccessToken      string
	RefreshToken     string
	RefreshExpiresIn int
	AccessExpiresIn  int
}
