package handler

import (
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

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAuthorization, accessToken string)

	tests := []struct {
		name                 string
		headerName           string
		headerValue          string
		accessToken          string
		mockBehavior         mockBehavior
		atUUIDKey            string
		atUUIDValue          string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			accessToken: "token",
			mockBehavior: func(r *mock_service.MockAuthorization, accessToken string) {
				r.EXPECT().VerifyUserToken(accessToken).Return(&app.AccessTokenClaims{
					UserID: 1,
					AtUUID: "atuuid",
				}, nil)
			},
			atUUIDKey:            "atuuid",
			atUUIDValue:          "valueATuuid",
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:                 "Empty Auth Header",
			headerName:           "Authorization",
			headerValue:          "",
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name:                 "Invalid Auth Header",
			headerName:           "Authorization",
			headerValue:          "invalid_header",
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Empty Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"token is empty"}`,
		},
		{
			name:        "Invalid Token",
			headerName:  "Authorization",
			headerValue: "Bearer invalid_token",
			accessToken: "invalid_token",
			mockBehavior: func(r *mock_service.MockAuthorization, accessToken string) {
				r.EXPECT().VerifyUserToken(accessToken).Return(nil, errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
		{
			name:        "Redis Error",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			accessToken: "token",
			mockBehavior: func(r *mock_service.MockAuthorization, accessToken string) {
				r.EXPECT().VerifyUserToken(accessToken).Return(&app.AccessTokenClaims{
					UserID: 1,
					AtUUID: "invalid uuid",
				}, nil)
			},
			atUUIDKey:            "atuuid",
			atUUIDValue:          "valueATuuid",
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"redis error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			if test.mockBehavior != nil {
				test.mockBehavior(repo, test.accessToken)
			}

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			r := gin.New()
			r.GET("/identity", handler.userIdentity)

			redis := app.GetRedisConn()
			redis.Set(context.Background(), test.atUUIDKey, test.atUUIDValue, 0)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
