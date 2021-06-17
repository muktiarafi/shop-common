package common

type UserPayload struct {
	ID    int      `json:"id"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

type SourceLocation struct {
	File     string
	Function string
	Line     int
}
