package service

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
	"time"
	"fmt"
	"crypto/sha1"
	"github.com/dgrijalva/jwt-go"
	"errors"
)

const (
	tokenTTLup = 10 * time.Minute
	tokenTTLdown = 1 * time.Second
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	salt = "hjqrhjqw124617ajfhajs"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
	UserLogin string `json:"login"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateAdmin(admin app.User) (int, error) {
	admin.Password = generatePasswordHash(admin.Password)
	
	return s.repo.CreateAdmin(admin)
}

func (s *AuthService) CreateUser(user app.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTLup).Unix(),
		IssuedAt: time.Now().Unix(),
		},
		UserId: user.Id,
		UserLogin: user.Login,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}
	
	return claims.UserId, claims.UserLogin, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}