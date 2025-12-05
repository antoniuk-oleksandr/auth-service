package errorhandler

import (
	"net/http"
	authDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/auth"
	usersDomain "github.com/antoniuk-oleksandr/auth-service/backend/internal/domain/users"
)

var errorStatusMap = map[error]int{
	// Users domain errors
	usersDomain.ErrUserNotFound:       http.StatusNotFound,
	usersDomain.ErrUserNotFound:       http.StatusNotFound,
	usersDomain.ErrUsernameTaken:      http.StatusConflict,
	usersDomain.ErrFailedToCreateUser: http.StatusInternalServerError,
	usersDomain.ErrFailedToFindUser:   http.StatusInternalServerError,

	//Auth domain errors
	authDomain.ErrInvalidCredentials: http.StatusUnauthorized,
}
