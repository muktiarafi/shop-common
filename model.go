package common

type UserPayload struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type SourceLocation struct {
	File     string
	Function string
	Line     int
}
