package app

type User struct {
	ID       int
	Login    string `json:"login" binding:"required"`
	Password string `json:"password_hash" binding:"required"`
	Age      string `json:"age" binding:"required"`
}

type UserDataInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password_hash" binding:"required"`
}

type RefreshDataInput struct {
	RefreshToken string `json:"refresh_token"`
}

type LogoutDataInput struct {
	AccessToken string `json:"access_token"`
}
