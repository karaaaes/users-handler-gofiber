package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"
	"go-fiber/model/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerRead(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"data": "Ini data user",
	})
}

func GetAllUser(ctx *fiber.Ctx) error {
	// Siapkan variable untuk menampung
	var users []entity.User

	// Panggil variable db
	result := database.DB.Find(&users).Error
	if result != nil {
		panic("Failed to find in database")
	}

	return ctx.JSON(fiber.Map{
		"data": users,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(request.UserRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.User{
		Name:     user.Name,
		Address:  user.Address,
		Phone:    user.Phone,
		Email:    user.Email,
		Password: user.Password,
	}

	errFindEmail := database.DB.Find(&newUser, "email = ?", user.Email).RowsAffected
	if errFindEmail > 0 {
		return ctx.JSON(fiber.Map{
			"message": "Email sudah ada",
		})
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.JSON(fiber.Map{
			"message": errCreateUser,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "data created",
		"data":    newUser,
	})
}

func GetUserById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	// Siapkan variable untuk menampung
	var user entity.User

	// Panggil variable db
	result := database.DB.Find(&user, "id = ?", userId)
	if result.RowsAffected == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Address:   user.Address,
		Phone:     user.Phone,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"data":    userResponse,
		"message": "success",
	})
}

func UpdateUserById(ctx *fiber.Ctx) error {
	UserUpdateRequest := new(request.UserUpdateRequest)

	if err := ctx.BodyParser(UserUpdateRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	userId := ctx.Params("id")
	// Siapkan variable untuk menampung
	var user entity.User

	// Panggil variabel DB
	result := database.DB.Find(&user, "id = ?", userId)
	if result.Error != nil {
		// Handle error dari database
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	if result.RowsAffected == 0 {
		// Jika tidak ada baris yang terpengaruh, user tidak ditemukan
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Update data user berdasarkan UserRequest yang tidak kosong
	if UserUpdateRequest.Name != "" {
		user.Name = UserUpdateRequest.Name
	}

	if UserUpdateRequest.Email != "" {
		user.Email = UserUpdateRequest.Email
	}

	if UserUpdateRequest.Address != "" {
		user.Address = UserUpdateRequest.Address
	}

	if UserUpdateRequest.Phone != "" {
		user.Phone = UserUpdateRequest.Phone
	}

	if UserUpdateRequest.Password != "" {
		user.Password = UserUpdateRequest.Password
	}

	user.UpdatedAt = time.Now()

	// Simpan perubahan ke database
	err := database.DB.Save(&user).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	return ctx.JSON(fiber.Map{
		"data":    user,
		"message": "success",
	})
}

func UpdatePasswordById(ctx *fiber.Ctx) error {
	UserUpdatePasswordRequest := new(request.UserUpdatePasswordRequest)

	if err := ctx.BodyParser(UserUpdatePasswordRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	userId := ctx.Params("id")
	var user entity.User

	result := database.DB.Find(&user, "id = ?", userId)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	if UserUpdatePasswordRequest.Password == "" {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Password belum diisi !",
		})
	}

	user.Password = UserUpdatePasswordRequest.Password

	// Simpan perubahan ke database
	if errUpdatePassword := database.DB.Save(&user).Error; errUpdatePassword != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to update password",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Password berhasil di update",
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	result := database.DB.First(&user, "id = ?", userId)
	if result.Error != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Not found",
		})
	}

	resultDelete := database.DB.Delete(&user).Error
	if resultDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}
