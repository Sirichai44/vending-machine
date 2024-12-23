package apis

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandleResponse(t *testing.T) {
	testCases := []struct {
		name          string
		configMock    func() fiber.Handler
		checkResponse func(t *testing.T, res *http.Response)
	}{
		{
			name: "Should be response success",
			configMock: func() fiber.Handler {
				return HandleResponse(func(c *fiber.Ctx) (any, error) {
					return fiber.Map{"code": "ok"}, nil
				})
			},
			checkResponse: func(t *testing.T, res *http.Response) {
				assert.Equal(t, fiber.StatusOK, res.StatusCode)
			},
		},
		{
			name: "Should be response error",
			configMock: func() fiber.Handler {
				return HandleResponse(func(c *fiber.Ctx) (any, error) {
					return nil, fiber.ErrBadRequest
				})
			},
			checkResponse: func(t *testing.T, res *http.Response) {
				assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
			},
		},
		{
			name: "Should be response slice nil",
			configMock: func() fiber.Handler {
				return HandleResponse(func(c *fiber.Ctx) (any, error) {
					var slice []string
					return slice, nil
				})
			},
			checkResponse: func(t *testing.T, res *http.Response) {
				assert.Equal(t, fiber.StatusOK, res.StatusCode)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := fiber.New()
			defer func() { _ = f.Shutdown() }()

			f.Get("/", tc.configMock())

			req := httptest.NewRequest(fiber.MethodGet, "/", nil)
			req.Header.Set("Content-Type", "application/json")

			res, err := f.Test(req)
			defer func() { _ = res.Body.Close() }()

			assert.Nil(t, err)
			tc.checkResponse(t, res)
		})
	}
}

func TestHandleBodyParser(t *testing.T) {
	type MyValidator struct {
		Name string `validate:"required"`
	}

	testCases := []struct {
		name          string
		request       string
		configMock    func() fiber.Handler
		checkResponse func(t *testing.T, res *http.Response)
	}{
		{
			name:    "Should be response success",
			request: `{"name": "body-parser"}`,
			configMock: func() fiber.Handler {
				return HandleBodyParser(func(req MyValidator, c *fiber.Ctx) (any, error) {
					return req, nil
				})
			},
			checkResponse: func(t *testing.T, res *http.Response) {
				assert.Equal(t, fiber.StatusOK, res.StatusCode)
			},
		},
		{
			name:    "Should be response error request logic error",
			request: `{"name": "body-parser"}`,
			configMock: func() fiber.Handler {
				return HandleBodyParser(func(req MyValidator, c *fiber.Ctx) (any, error) {
					return req, errors.New("logic return error")
				})
			},
			checkResponse: func(t *testing.T, res *http.Response) {
				assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
			},
		},
		{
			name:    "Should be response error request Validator",
			request: `{"name": ""}`,
			configMock: func() fiber.Handler {
				return HandleBodyParser(func(req MyValidator, c *fiber.Ctx) (any, error) {
					return req, nil
				})
			},
			checkResponse: func(t *testing.T, res *http.Response) {
				assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
			},
		},
		{
			name:    "Should be response error request NotAcceptable",
			request: `{"name": 0}`,
			configMock: func() fiber.Handler {
				return HandleBodyParser(func(req MyValidator, c *fiber.Ctx) (any, error) {
					return req, nil
				})
			},
			checkResponse: func(t *testing.T, res *http.Response) {
				assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			f := fiber.New()
			defer func() { _ = f.Shutdown() }()

			f.Post("/", tc.configMock())

			req := httptest.NewRequest(fiber.MethodPost, "/", strings.NewReader(tc.request))
			req.Header.Set("Content-Type", "application/json")

			res, err := f.Test(req)
			defer func() { _ = res.Body.Close() }()

			assert.Nil(t, err)
			tc.checkResponse(t, res)
		})
	}
}
