package users

type CreateUserCommand struct {
	Username     string
	PasswordHash string
}
