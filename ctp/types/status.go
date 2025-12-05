package types

type Status int32

const (
	StatusOK            Status = 200
	StatusCreated       Status = 201
	StatusBadRequest    Status = 400
	StatusUnauthorized  Status = 401
	StatusNotFound      Status = 404
	StatusInternalError Status = 500
)

var statusText = map[Status]string{
	StatusOK:            "OK",
	StatusCreated:       "Created",
	StatusBadRequest:    "Bad Request",
	StatusNotFound:      "Not Found",
	StatusInternalError: "Internal Server Error",
}

func (s Status) String() string {
	if text, exists := statusText[s]; exists {
		return text
	}
	return "Unknown Status"
}
