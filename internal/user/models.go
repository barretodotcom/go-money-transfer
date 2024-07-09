package user

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
