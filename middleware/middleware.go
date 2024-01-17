package middleware

import "github.com/gofiber/fiber/v2"

func UserMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token != "tokenauthorization" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}
