package apis

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"daveslist/apis/middleware"
// 	"daveslist/dtos"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// func TestGetListAll(t *testing.T) {
// 	contextPath := "/" + apiVersion + "/product"

// 	testCases := []struct {
// 		name          string
// 		configMock    func(*mock_service.MockProductService)
// 		checkResponse func(t *testing.T, res *http.Response)
// 	}{{
// 		name: "Should be response success",
// 		configMock: func(srvProduct *mock_service.MockProductService) {
// 			srvProduct.EXPECT().FindAll(gomock.Any()).Return([]dtos.Product{}, nil).Times(1)
// 		},

// 		checkResponse: func(t *testing.T, res *http.Response) {
// 			assert.Equal(t, fiber.StatusOK, res.StatusCode)
// 		},
// 	}}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			srvProduct := mock_service.NewMockProductService(ctrl)
// 			tc.configMock(srvProduct)

// 			f := fiber.New()
// 			f.Use(func(c *fiber.Ctx) error {
// 				c.Locals(dtos.KEY_AUTH_STATUS, true)
// 				c.Locals(dtos.KEY_SECRET, "test_sign_key")
// 				return c.Next()
// 			})

// 			defer func() { _ = f.Shutdown() }()

// 			NewHandlerProduct(f, srvProduct)
// 		})
// 	}
// }

// func TestGetListAll(t *testing.T) {
// 	contextPath := "/" + apiVersion + "/list"

// 	testCases := []struct {
// 		name          string
// 		configMock    func(*mock_service.MockListService)
// 		checkResponse func(t *testing.T, res *http.Response)
// 	}{{
// 		name: "Should be response success",
// 		configMock: func(srvList *mock_service.MockListService) {
// 			srvList.EXPECT().FindAll(gomock.Any()).Return([]dtos.Listing{}, nil).Times(1)
// 		},

// 		checkResponse: func(t *testing.T, res *http.Response) {
// 			assert.Equal(t, fiber.StatusOK, res.StatusCode)
// 		},
// 	}}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			srvAuth := mock_service.NewMockAuthService(ctrl)
// 			srvList := mock_service.NewMockListService(ctrl)
// 			tc.configMock(srvList)

// 			f := fiber.New()
// 			f.Use(func(c *fiber.Ctx) error {
// 				c.Locals(dtos.KEY_AUTH_STATUS, true)
// 				c.Locals(dtos.KEY_SECRET, "test_sign_key")
// 				return c.Next()
// 			})

// 			f.Use(middleware.CheckAuthMiddleWare)
// 			defer func() { _ = f.Shutdown() }()

// 			NewHandlerList(f, srvList, srvAuth)

// 			token, err := genToken("test", "test@mail.com", "test_sign_key")
// 			assert.Nil(t, err)

// 			req := httptest.NewRequest(http.MethodGet, contextPath, nil)
// 			req.Header.Set("Authorization", token)
// 			req.Header.Set("Content-Type", "application/json")

// 			res, err := f.Test(req, -1)
// 			defer func() { _ = res.Body.Close() }()

// 			assert.Nil(t, err)
// 			tc.checkResponse(t, res)
// 		})
// 	}
// }
