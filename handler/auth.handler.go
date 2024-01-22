package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"go-fiber/model/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func AuthLogin(ctx *fiber.Ctx) error {
	userRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// Check Email
	var user entity.User
	errFindEmail := database.DB.First(&user, "email = ?", userRequest.Email).Error
	if errFindEmail != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Email not found",
		})
	}

	// Check Password
	isValid := CheckHashPassword(userRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	claims := jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
	}

	token, err := utils.GenerateJWT(&claims)
	if err != nil {
		return ctx.Status(fiber.ErrBadGateway.Code).JSON(fiber.Map{
			"message": "internal server error JWT",
		})

	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

func CheckHashPassword(password, hash_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password))
	return err == nil
}
