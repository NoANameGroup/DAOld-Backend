package session

type CreateSessionReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
