package auth

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterDTO struct {
	Username    string `json:"username" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required"`
}
