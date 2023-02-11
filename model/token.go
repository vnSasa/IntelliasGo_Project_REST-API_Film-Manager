package app

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserID  int    `json:"user_id"`
	AtUUID  string `json:"access_uuid"`
	RtUUID  string `json:"refresh_uuid"`
	IsAdmin bool   `json:"is_admin"`
	IsUser  bool   `json:"is_user"`
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
	UserID    int    `json:"user_id"`
	RtUUID    string `json:"refresh_uuid"`
	AtUUID    string `json:"access_uuid"`
	IsAdmin   bool   `json:"is_admin"`
	IsUser    bool   `json:"is_user"`
	IsRefresh bool   `json:"is_refresh"`
}
