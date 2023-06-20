package dto

type UserRequest struct {
	Name     string `json:"nombre"`
	Email    string `json:"email"`
	Password string `json:"contrasena"`
	Role     string `json:"rol"`
}
