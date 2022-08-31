package models

type AuthResponse struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    AuthResponseData `json:"data"`
}

type AuthResponseData struct {
	Token string `json:"token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
}

type AuthFailedResponse struct {
	Status  int    `json:"status" example:"0"`
	Message string `json:"message" example:"Authentication failed"`
}

type TokenRequest struct {
	Email    string `json:"email" example:"user_repository@email.com"`
	Password string `json:"password" example:"your password"`
}
