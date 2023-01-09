package app

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenDetails struct {
	AccessToken string
	AccessUuid string
	AtExpires int64
} 

type TokenClaims struct {
	jwt.StandardClaims
	UserID  int    `json:"user_id"`
	AtUuid string `json:"access_uuid"`
	IsAdmin bool `json:"is_admin"`
	IsUser bool `json:"is_user"`
}