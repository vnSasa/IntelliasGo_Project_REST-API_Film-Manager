package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

const (
	tokenTTLup          = 10 * time.Minute
	signingKey          = "qrkjk#4#%35FSFJlja#4353KSFjH"
	salt                = "hjqrhjqw124617ajfhajs"
	timeForAccessToken  = 15
	timeForRefreshToken = 24 * 7
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
		isUser  = true
	)

	if strings.Compare(login, os.Getenv("ADMIN_LOGIN")) == 0 {
		isAdmin = true
		isUser = false
	}

	td := &app.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * timeForAccessToken).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * timeForRefreshToken).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &app.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.AtExpires,
		},
		UserID:  user.ID,
		AtUUID:  td.AccessUUID,
		RtUUID:  td.RefreshUUID,
		IsAdmin: isAdmin,
		IsUser:  isUser,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &app.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.RtExpires,
		},
		UserID:    user.ID,
		RtUUID:    td.RefreshUUID,
		AtUUID:    td.AccessUUID,
		IsAdmin:   isAdmin,
		IsUser:    isUser,
		IsRefresh: true,
	})

	td.AccessToken, err = accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	td.RefreshToken, err = refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *AuthService) RefreshToken(refreshData *app.RefreshTokenClaims) (*app.TokenDetails, error) {
	td := &app.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * timeForAccessToken).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * timeForRefreshToken).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &app.AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.AtExpires,
		},
		UserID:  refreshData.UserID,
		AtUUID:  td.AccessUUID,
		RtUUID:  td.RefreshUUID,
		IsAdmin: refreshData.IsAdmin,
		IsUser:  refreshData.IsUser,
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &app.RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.RtExpires,
		},
		UserID:    refreshData.UserID,
		RtUUID:    td.RefreshUUID,
		AtUUID:    td.AccessUUID,
		IsAdmin:   refreshData.IsAdmin,
		IsUser:    refreshData.IsUser,
		IsRefresh: true,
	})
	var err error
	td.AccessToken, err = accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}
	td.RefreshToken, err = refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (s *AuthService) ParseToken(accessToken string) (*app.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &app.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*app.AccessTokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims, nil
}

func (s *AuthService) VerifyAdminToken(accessToken string) (*app.AccessTokenClaims, error) {
	claims, err := s.ParseToken(accessToken)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if !claims.IsAdmin {
		return nil, errors.New("only admin have access")
	}

	return claims, nil
}

func (s *AuthService) VerifyUserToken(accessToken string) (*app.AccessTokenClaims, error) {
	claims, err := s.ParseToken(accessToken)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if !claims.IsUser {
		return nil, errors.New("only user have access")
	}

	return claims, nil
}

func (s *AuthService) ParseRefreshToken(refreshToken string) (*app.RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &app.RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*app.RefreshTokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	if !claims.IsRefresh {
		return nil, errors.New("is not refresh token")
	}

	return claims, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
