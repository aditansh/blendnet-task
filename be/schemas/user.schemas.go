package schemas

type RegisterUserSchema struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserSchema struct {
	Name string `json:"name,omitempty" validate:"required,min=3"`
}

type UpdateWatchlistSchema struct {
	Watchlist []string `json:"watchlist" validate:"required"`
}
