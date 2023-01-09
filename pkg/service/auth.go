package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
	"strings"
	"os"

	"github.com/twinj/uuid"
	"github.com/dgrijalva/jwt-go"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

const (
	tokenTTLup = 10 * time.Minute
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	salt       = "hjqrhjqw124617ajfhajs"
)

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

func (s *AuthService) GenerateToken(login, password string) (*app.TokenDetails, error) {
	user, err := s.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return nil, err
	}

	var (
		isAdmin = false
		isUser = true
	)

	if strings.Compare(login, os.Getenv("ADMIN_LOGIN")) == 0 {
		isAdmin = true
		isUser = false
	}

	td := &app.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &app.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.AtExpires,
		},
		UserID:  user.ID,
		AtUuid: td.AccessUuid,
		IsAdmin: isAdmin,
		IsUser: isUser,
	})

	td.AccessToken, err = token.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *AuthService) ParseToken(accessToken string) (*app.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &app.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*app.TokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
