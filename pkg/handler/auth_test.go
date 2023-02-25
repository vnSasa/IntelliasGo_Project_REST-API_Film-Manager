package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"

	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/service"
	mock_service "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, user app.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            app.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"login": "login", "password_hash": "qwerty", "age": "testage"}`,
			inputUser: app.User{
				Login:    "login",
				Password: "qwerty",
				Age:      "testage",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user app.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"login": "login"}`,
			inputUser:            app.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user app.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"login": "login", "password_hash": "qwerty", "age": "testage"}`,
			inputUser: app.User{
				Login:    "login",
				Password: "qwerty",
				Age:      "testage",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user app.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, user app.UserDataInput)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            app.UserDataInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"login": "login", "password_hash": "qwerty"}`,
			inputUser: app.UserDataInput{
				Login:    "login",
				Password: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user app.UserDataInput) {
				r.EXPECT().GenerateToken(user.Login, user.Password).
					Return(&app.TokenDetails{
						AccessToken:  "Atoken",
						RefreshToken: "Rtoken",
					}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"accsess_token":"Atoken","refresh_token":"Rtoken"}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{}`,
			inputUser:            app.UserDataInput{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user app.UserDataInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"login": "login", "password_hash": "qwerty"}`,
			inputUser: app.UserDataInput{
				Login:    "login",
				Password: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user app.UserDataInput) {
				r.EXPECT().GenerateToken(user.Login, user.Password).
					Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/sign-in", handler.signIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_refreshSignIn(t *testing.T) {
	type mockBehaviorData func(r *mock_service.MockAuthorization, data app.RefreshDataInput)
	type mockBehaviorClaims func(r *mock_service.MockAuthorization, input app.RefreshTokenClaims)

	tests := []struct {
		name                 string
		inputBody            string
		inputData            app.RefreshDataInput
		mockBehaviorData     mockBehaviorData
		atUUIDKey            string
		atUUIDValue          string
		rtUUIDKey            string
		rtUUIDValue          string
		inputClaims          app.RefreshTokenClaims
		mockBehaviorClaims   mockBehaviorClaims
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"refresh_token": "refresh_token"}`,
			inputData: app.RefreshDataInput{
				RefreshToken: "refresh_token",
			},
			mockBehaviorData: func(r *mock_service.MockAuthorization, data app.RefreshDataInput) {
				r.EXPECT().ParseRefreshToken(data.RefreshToken).Return(&app.RefreshTokenClaims{
					AtUUID:    "atkey",
					RtUUID:    "rtkey",
					IsRefresh: true,
				}, nil)
			},
			rtUUIDKey:   "rtkey",
			rtUUIDValue: "rtvalue",
			inputClaims: app.RefreshTokenClaims{
				AtUUID:    "atkey",
				RtUUID:    "rtkey",
				IsRefresh: true,
			},
			mockBehaviorClaims: func(r *mock_service.MockAuthorization, input app.RefreshTokenClaims) {
				r.EXPECT().RefreshToken(&input).Return(&app.TokenDetails{
					AccessToken:  "Atoken",
					RefreshToken: "Rtoken",
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"access_token":"Atoken","refresh_token":"Rtoken"}`,
		},
		{
			name:      "Invalid Refresh Token",
			inputBody: `{"refresh_token": "invalid_refresh_token"}`,
			inputData: app.RefreshDataInput{
				RefreshToken: "invalid_refresh_token",
			},
			mockBehaviorData: func(r *mock_service.MockAuthorization, data app.RefreshDataInput) {
				r.EXPECT().ParseRefreshToken(data.RefreshToken).Return(nil, errors.New("invalid refresh token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid refresh token"}`,
		},
		{
			name:      "Redis Error",
			inputBody: `{"refresh_token": "refresh_token"}`,
			inputData: app.RefreshDataInput{
				RefreshToken: "refresh_token",
			},
			mockBehaviorData: func(r *mock_service.MockAuthorization, data app.RefreshDataInput) {
				r.EXPECT().ParseRefreshToken(data.RefreshToken).Return(&app.RefreshTokenClaims{
					RtUUID:    "invalid_rtkey",
					IsRefresh: true,
				}, nil)
			},
			rtUUIDKey:            "rtkey",
			rtUUIDValue:          "rtvalue",
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"redis error"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"refresh_token": "refresh_token"}`,
			inputData: app.RefreshDataInput{
				RefreshToken: "refresh_token",
			},
			mockBehaviorData: func(r *mock_service.MockAuthorization, data app.RefreshDataInput) {
				r.EXPECT().ParseRefreshToken(data.RefreshToken).Return(&app.RefreshTokenClaims{
					AtUUID:    "atkey",
					RtUUID:    "rtkey",
					IsRefresh: true,
				}, nil)
			},
			rtUUIDKey:   "rtkey",
			rtUUIDValue: "rtvalue",
			inputClaims: app.RefreshTokenClaims{
				AtUUID:    "atkey",
				RtUUID:    "rtkey",
				IsRefresh: true,
			},
			mockBehaviorClaims: func(r *mock_service.MockAuthorization, input app.RefreshTokenClaims) {
				r.EXPECT().RefreshToken(&input).Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			if test.mockBehaviorData != nil {
				test.mockBehaviorData(repo, test.inputData)
			}
			if test.mockBehaviorClaims != nil {
				test.mockBehaviorClaims(repo, test.inputClaims)
			}

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/refresh", handler.refreshSignIn)

			redis := app.GetRedisConn()
			redis.Set(context.Background(), test.rtUUIDKey, test.rtUUIDValue, 0)
			redis.Set(context.Background(), test.atUUIDKey, test.atUUIDValue, 0)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/refresh",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_logout(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, token app.LogoutDataInput)

	tests := []struct {
		name                 string
		inputBody            string
		inputData            app.LogoutDataInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"access_token": "token"}`,
			inputData: app.LogoutDataInput{
				AccessToken: "token",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, token app.LogoutDataInput) {
				r.EXPECT().ParseToken(token.AccessToken).Return(&app.AccessTokenClaims{}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `"Logout Success"`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"access_token"}`,
			inputData:            app.LogoutDataInput{},
			mockBehavior:         func(r *mock_service.MockAuthorization, token app.LogoutDataInput) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"access_token": "token"}`,
			inputData: app.LogoutDataInput{
				AccessToken: "token",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, token app.LogoutDataInput) {
				r.EXPECT().ParseToken(token.AccessToken).Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputData)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/logout", handler.logout)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/logout",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
