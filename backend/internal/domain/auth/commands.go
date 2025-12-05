package auth

type (
	RegisterCommand struct {
		Username   string
		Password   string
	}

	LoginCommand struct {
		Username string
		Password string
	}
)
