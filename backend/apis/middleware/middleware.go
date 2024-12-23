package middleware

// import (
// 	"errors"
// 	"log"
// 	"time"

// 	"vending_machine/dtos"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v5"
// )

// func CheckAuthMiddleWare(c *fiber.Ctx) error {
// 	var authen bool
// 	secretKey := c.Locals(dtos.KEY_SECRET)

// 	token := c.Get("Authorization")
// 	if token == "" {
// 		authen = false
// 	} else {
// 		user, err := getUserFromToken(token, secretKey.(string))
// 		if err != nil {
// 			authen = false
// 		} else {
// 			c.Locals("user", user)
// 			authen = true
// 		}
// 	}

// 	c.Locals(dtos.KEY_AUTH_STATUS, authen)

// 	return c.Next()
// }

// func IsAuthenticated(c *fiber.Ctx) error {
// 	secretKey := c.Locals(dtos.KEY_SECRET).(string)

// 	token := c.Get("Authorization")
// 	if token == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No token provided"})
// 	}

// 	user, err := getUserFromToken(token, secretKey)
// 	if err != nil {
// 		log.Println("IsAuthenticated", "getUserFromToken", "failed")
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	c.Locals(dtos.KEY_USER, user)
// 	return c.Next()
// }

// func getUserFromToken(token string, scKey string) (*dtos.User, error) {
// 	claims := &dtos.Claims{}
// 	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fiber.ErrUnauthorized
// 		}
// 		return []byte(scKey), nil
// 	})

// 	if err != nil || !tkn.Valid {
// 		return nil, err
// 	}

// 	expire := claims.ExpiresAt.Time

// 	if expire.Before(time.Now()) {
// 		return nil, errors.New("token is expired")
// 	}

// 	user := dtos.User{
// 		Email:    claims.Email,
// 		Username: claims.Username,
// 	}

// 	return &user, nil
// }
