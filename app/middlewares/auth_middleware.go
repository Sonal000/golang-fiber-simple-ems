package middleware

import (
	"employeeManagement/app/dto"
	"employeeManagement/app/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthProtect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtService := services.NewJWTService()

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.Response{
				Message: "unauthorized access",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.Response{
				Message: "Invalid or expired token",
			})
		}
		c.Locals("userClaims", claims)
		return c.Next()
	}
}

func AuthRestrict(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims, ok := c.Locals("userClaims").(map[string]interface{})
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(dto.Response{
				Message: "Access forbidden",
			})
		}
		role, ok := userClaims["role"].(string)
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(dto.Response{
				Message: "Role not found or invalid role type",
			})
		}
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(dto.Response{
			Message: "You do not have permission to access this resource",
		})
	}
}
