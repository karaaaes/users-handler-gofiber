package middleware

import (
	"fmt"
	"go-fiber/model/utils"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	_, err := utils.VerifyJWT(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil {
		// Handle error
		fmt.Println("Error decoding token:", err)
	} else {
		// Gunakan nilai claims yang telah didapatkan
		fmt.Println("Decoded Claims:", claims)
	}

	return ctx.Next()
}
