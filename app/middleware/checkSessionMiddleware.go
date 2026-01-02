package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/core-wallet/app/client"
	"github.com/core-wallet/app/config"
	"github.com/core-wallet/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"gorm.io/gorm"
)

func SessionAuthentication(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			utils.ErrorLog("MIDDLEWARE", errors.New("token required"), false)
			return utils.ResponseError(c, fiber.StatusUnauthorized, "missing authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Bypass untuk app secret
		if tokenStr == config.AppConfig.AppSecret {
			return c.Next()
		}

		// Parse dan validasi JWT
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.AppConfig.AppSecret), nil
		})
		if err != nil || !token.Valid {
			utils.ErrorLog("MIDDLEWARE", fmt.Errorf("invalid token: %w", err), false)
			return utils.ResponseError(c, fiber.StatusUnauthorized, "invalid token")
		}

		// Validasi user session
		user, err := client.GetUserSession(tokenStr, "session_authentication")
		if err != nil {
			utils.ErrorLog("MIDDLEWARE", fmt.Errorf("invalid session: %w", err), false)
			return utils.ResponseError(c, fiber.StatusUnauthorized, "invalid token session")
		}

		utils.InfoLog(fmt.Sprintf("User authenticated - ID: %v, Email: %s", user.ID, user.Email))

		c.Locals("user_id", user.ID)
		c.Locals("email", user.Email)
		c.Locals("token", tokenStr)

		return c.Next()
	}
}
