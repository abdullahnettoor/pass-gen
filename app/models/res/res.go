package res

type User struct {
	UserID   string
	UserName string
}

type LoginResponse struct {
	UserID   string
	UserName string
	Password string
}
