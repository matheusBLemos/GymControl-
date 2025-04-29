package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

// Middleware de autenticação para proteger rotas
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header missing",
			})
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("DB_NAME")), nil // usa o mesmo secret da geração do token
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Armazena os dados do token no contexto
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("user_id", claims["user_id"])
			c.Locals("name", claims["name"])
			c.Locals("email", claims["email"])
		}

		return c.Next()
	}
}

// Middleware utilitário para retornar os dados do token
func GetTokenData(c *fiber.Ctx) map[string]interface{} {
	return map[string]interface{}{
		"user_id": c.Locals("user_id"),
		"name":    c.Locals("name"),
		"email":   c.Locals("email"),
	}
}
