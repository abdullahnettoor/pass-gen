package res

type User struct {
	ID       string
	UserName string
}

type LoginResponse struct {
	ID       string
	UserName string
	Password string
}

type SecretsCollectonResponse struct {
	Name []string
}

type SecretResponse struct {
	Name            string
	Secret          []byte
	SecretPlainText string
}
