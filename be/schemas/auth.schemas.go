package schemas

type LoginSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponseSchema struct {
	RefreshToken string `json:"refreshToken"`
}

type UpdatePasswordSchema struct {
	OldPassword string `json:"oldPassword" validate:"required,min=6"`
	NewPassword string `json:"newPassword" validate:"required,min=6"`
}

type RefreshTokenSchema struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type LogoutSchema struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
