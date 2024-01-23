package handler

import (
	"go-fiber/database"
	"go-fiber/model/entity"
	"go-fiber/model/request"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func CreateComic(ctx *fiber.Ctx) error {
	comic := new(request.ComicRequest)
	if err := ctx.BodyParser(comic); err != nil {
		return err
	}

	// Validate Request
	validate := validator.New()
	errValidate := validate.Struct(comic)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Dapatkan nama file dari konteks
	fileName, ok := ctx.Locals("fileName").(string)
	if !ok || fileName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File name not found",
		})
	}

	newComic := entity.Comic{
		SeriesName: comic.SeriesName,
		Author:     comic.Author,
		Cover:      fileName,
	}

	//  UPLOAD MULTIPLE FILES
	// // Dapatkan nama-nama file dari konteks
	// fileNames, ok := ctx.Locals("fileNames").([]string)
	// if !ok || len(fileNames) == 0 {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "File names not found",
	// 	})
	// }

	// // Gabungkan semua nama file menjadi satu string
	// combinedNames := strings.Join(fileNames, ", ")

	// newComic := entity.Comic{
	// 	SeriesName: comic.SeriesName,
	// 	Author:     comic.Author,
	// 	Cover:      combinedNames,
	// }

	errFindEmail := database.DB.Find(&newComic, "series_name = ?", comic.SeriesName).RowsAffected
	if errFindEmail > 0 {
		return ctx.JSON(fiber.Map{
			"message": "Comic already created",
		})
	}

	errCreateComic := database.DB.Create(&newComic).Error
	if errCreateComic != nil {
		return ctx.JSON(fiber.Map{
			"message": errCreateComic,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "comic created",
		"data":    newComic,
	})
}
