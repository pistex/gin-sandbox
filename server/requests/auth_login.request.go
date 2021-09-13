package requests

type AuthLogin struct {
	Email             string `json:"email" binding:"required,email"`
	Password          string `json:"password" binding:"required,min=8"`
	AuthorizationCode string `json:"authorizationCode" binding:"required"`
	CodeVerifier      string `json:"codeVerifier" binding:"required"`
}
