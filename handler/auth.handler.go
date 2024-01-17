package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"

	"github.com/gofiber/fiber/v2"
)

func AuthLogin(ctx *fiber.Ctx) error {
	userRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	var user entity.User
	errFindEmail := database.DB.First(&user, "email = ?", userRequest.Email).Error
	if errFindEmail != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Email not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "tokenauthorization",
	})
}
