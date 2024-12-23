package middleware

// import (
// 	"net/http/httptest"
// 	"testing"

// 	"vending_machine/dtos"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCheckAuthMiddleWare(t *testing.T) {
// 	secretKey := "test_secret_key"

// 	f := fiber.New()
// 	f.Use(func(c *fiber.Ctx) error {
// 		c.Locals(dtos.KEY_SECRET, secretKey)
// 		return c.Next()
// 	})
// 	f.Use(CheckAuthMiddleWare)

// 	// Define test cases
// 	testCases := []struct {
// 		name           string
// 		token          string
// 		expectedStatus int
// 	}{
// 		{
// 			name:           "No token provided",
// 			token:          "",
// 			expectedStatus: fiber.StatusNotFound,
// 		},
// 		{
// 			name:           "Invalid token",
// 			token:          "invalid_token",
// 			expectedStatus: fiber.StatusNotFound,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			req := httptest.NewRequest("GET", "/", nil)
// 			req.Header.Set("Authorization", tc.token)

// 			res, err := f.Test(req)
// 			defer func() { _ = res.Body.Close() }()

// 			assert.Nil(t, err)
// 			assert.Equal(t, tc.expectedStatus, res.StatusCode)
// 		})
// 	}
// }

// func TestIsAuthenticated(t *testing.T) {
// 	secretKey := "test_secret_key"

// 	f := fiber.New()
// 	f.Use(func(c *fiber.Ctx) error {
// 		c.Locals(dtos.KEY_SECRET, secretKey)
// 		return c.Next()
// 	})
// 	f.Use(IsAuthenticated)

// 	testCases := []struct {
// 		name           string
// 		token          string
// 		expectedStatus int
// 	}{
// 		{
// 			name:           "No token provided",
// 			token:          "",
// 			expectedStatus: fiber.StatusUnauthorized,
// 		},
// 		{
// 			name:           "Invalid token",
// 			token:          "invalid_token",
// 			expectedStatus: fiber.StatusUnauthorized,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			req := httptest.NewRequest("GET", "/", nil)
// 			req.Header.Set("Authorization", tc.token)

// 			res, err := f.Test(req)
// 			defer func() { _ = res.Body.Close() }()

// 			assert.Nil(t, err)
// 			assert.Equal(t, tc.expectedStatus, res.StatusCode)
// 		})
// 	}
// }
