package models

type UserCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Company  struct {
		Nome string `json:"nome"`
		CNPJ string `json:"cnpj"`
	} `json:"company"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
