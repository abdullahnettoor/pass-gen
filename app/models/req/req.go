package req

type User struct {
	UserName        string
	Password        string
	ConfirmPassword string
}

type Token struct {
	Token string `json:"token"`
}
