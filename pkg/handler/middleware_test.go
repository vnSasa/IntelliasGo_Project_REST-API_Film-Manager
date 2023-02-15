package handler

// import (
// 	"fmt"
// 	// "errors"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/magiconair/properties/assert"
// 	// "github.com/stretchr/testify/assert"
// 	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
// 	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/service"
// 	mock_service "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/service/mocks"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestHandler_parseAuthHeader(t *testing.T) {
// 	type mockBehavior func(r *mock_service.MockAuthorization, token string)

// 	accessClaims := app.AccessTokenClaims{}

// 	tests := []struct {
// 		name	string
// 		headerName	string
// 		headerValue	string
// 		token	string
// 		mockBehavior	mockBehavior
// 		expectedStatusCode	int
// 		expectedResponseBody	app.AccessTokenClaims
// 	}{
// 		{
// 			name:        "Ok",
// 			headerName:  "Authorization",
// 			headerValue: "Bearer token",
// 			token:       "token",
// 			mockBehavior: func(r *mock_service.MockAuthorization, token string) {
// 				r.EXPECT().ParseToken(token).Return(&accessClaims{

// 				}, nil)
// 			},
// 			expectedStatusCode:   200,
// 			expectedResponseBody: accessClaims{

// 			},
// 		},
// 		// {
// 		// 	name:                 "Invalid Header Name",
// 		// 	headerName:           "",
// 		// 	headerValue:          "Bearer token",
// 		// 	token:                "token",
// 		// 	mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
// 		// 	expectedStatusCode:   401,
// 		// 	expectedResponseBody: `{"message":"empty auth header"}`,
// 		// },
// 		// {
// 		// 	name:                 "Invalid Header Value",
// 		// 	headerName:           "Authorization",
// 		// 	headerValue:          "Bearr token",
// 		// 	token:                "token",
// 		// 	mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
// 		// 	expectedStatusCode:   401,
// 		// 	expectedResponseBody: `{"message":"invalid auth header"}`,
// 		// },
// 		// {
// 		// 	name:                 "Empty Token",
// 		// 	headerName:           "Authorization",
// 		// 	headerValue:          "Bearer ",
// 		// 	token:                "token",
// 		// 	mockBehavior:         func(r *mock_service.MockAuthorization, token string) {},
// 		// 	expectedStatusCode:   401,
// 		// 	expectedResponseBody: `{"message":"token is empty"}`,
// 		// },
// 		// {
// 		// 	name:        "Parse Error",
// 		// 	headerName:  "Authorization",
// 		// 	headerValue: "Bearer token",
// 		// 	token:       "token",
// 		// 	mockBehavior: func(r *mock_service.MockAuthorization, token string) {
// 		// 		r.EXPECT().ParseToken(token).Return(0, errors.New("invalid token"))
// 		// 	},
// 		// 	expectedStatusCode:   401,
// 		// 	expectedResponseBody: `{"message":"invalid token"}`,
// 		// },
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			repo := mock_service.NewMockAuthorization(c)
// 			test.mockBehavior(repo, test.token)

// 			services := &service.Service{Authorization: repo}
// 			handler := Handler{services}

// 			r := gin.New()
// 			r.GET("/identity", handler.userIdentity, func(c *gin.Context) {
// 				id, _ := c.Get(userCtx)
// 				c.String(200, fmt.Sprintf("%d", id.(int)))
// 			})

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest("GET", "/identity", nil)
// 			req.Header.Set(test.headerName, test.headerValue)

// 			r.ServeHTTP(w, req)

// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
// }

// // func TestGetUserId(t *testing.T) {
// // 	var getContext = func(id int) *gin.Context {
// // 		ctx := &gin.Context{}
// // 		ctx.Set(userCtx, id)
// // 		return ctx
// // 	}

// // 	tests := []struct {
// // 		name       string
// // 		ctx        *gin.Context
// // 		id         int
// // 		shouldFail bool
// // 	}{
// // 		{
// // 			name: "Ok",
// // 			ctx:  getContext(1),
// // 			id:   1,
// // 		},
// // 		{
// // 			ctx:        &gin.Context{},
// // 			name:       "Empty",
// // 			shouldFail: true,
// // 		},
// // 	}

// // 	for _, test := range tests {
// // 		t.Run(test.name, func(t *testing.T) {
// // 			id, err := getUserId(test.ctx)
// // 			if test.shouldFail {
// // 				assert.Error(t, err)
// // 			} else {
// // 				assert.NoError(t, err)
// // 			}

// // 			assert.Equal(t, id, test.id)
// // 		})
// // 	}
// // }
