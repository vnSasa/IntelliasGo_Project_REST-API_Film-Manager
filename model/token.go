package app

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenDetails struct {
	AccessToken string
	RefreshToken string
	AccessUuid string
	RefreshUuid string
	AtExpires int64
	RtExpires int64
} 

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserID  int    `json:"user_id"`
	AtUuid string `json:"access_uuid"`
	IsAdmin bool `json:"is_admin"`
	IsUser bool `json:"is_user"`
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
	UserID  int    `json:"user_id"`
	RtUuid string `json:"refresh_uuid"`
	IsAdmin bool `json:"is_admin"`
	IsUser bool `json:"is_user"`
}