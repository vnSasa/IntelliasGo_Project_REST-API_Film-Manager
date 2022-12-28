package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

const (
	tokenTTLup = 10 * time.Minute
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	salt       = "hjqrhjqw124617ajfhajs"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID  int    `json:"user_id"`
	UserAge string `json:"age"`
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

func (s *AuthService) GetLoginByID(id int) (string, error) {
	return s.repo.GetLoginByID(id)
}

func (s *AuthService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

func (s *AuthService) GetUser(login, password string) error {
	_, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTLup).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID:  user.ID,
		UserAge: user.Age,
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

	return claims.UserID, claims.UserAge, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
