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

type SecretsCollectionResponse struct {
	Name []string
}

type SecretPasswordResponse struct {
	Key             string
	Password        []byte
	SecretPlainText string
}
