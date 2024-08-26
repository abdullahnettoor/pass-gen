package req

type User struct {
	UserName        string
	Password        string
	ConfirmPassword string
}

type Token struct {
	Token string `json:"token"`
}

type Credential struct {
	UserID     string
	Key        string
	Secret     string
	CipherText []byte
}

type GetSecretPassword struct {
	UserID string
	Key    string
}

type GetKey struct {
	UserID string
}
