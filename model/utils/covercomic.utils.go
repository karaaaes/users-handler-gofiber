package utils

import (
	"fmt"
	"go-fiber/model/request"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func HandleCoverComic(ctx *fiber.Ctx) error {
	comic := new(request.ComicRequest)
	if err := ctx.BodyParser(comic); err != nil {
		return err
	}

	// Validate File
	fileCover, err := ctx.FormFile("cover")
	if err != nil {
		fmt.Println("Error : ", err)
		return err
	}

	var fileName string
	var extensionFile string

	if fileCover != nil {
		// Ambil ekstensi dari fileCover.Filename
		extensionFile = filepath.Ext(fileCover.Filename)
		fmt.Println(extensionFile)

		// Tentukan logika penamaan file berdasarkan kondisi
		if comic.SeriesName != "" {
			// Jika SeriesName tidak kosong, gunakan CoverName
			fileName = comic.CoverName + extensionFile
		} else {
			// Jika SeriesName kosong, gunakan fileCover.Filename
			fileName = fileCover.Filename
		}

		errSaveFile := ctx.SaveFile(fileCover, "./public/asset/cover/"+fileName)
		if errSaveFile != nil {
			fmt.Println("Error : ", errSaveFile)
			return errSaveFile
		}
	} else {
		fmt.Println("No File Uploaded...")
	}

	// Simpan nama file dan ekstensi dalam konteks Fiber
	fmt.Println(fileName)
	ctx.Locals("fileName", fileName)

	return ctx.Next()
}

func HandleMultipleCovers(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	var fileNames []string

	for _, files := range form.File {
		for _, file := range files {
			// Generate unique filename
			fileName := filepath.Join("./public/asset/cover/", file.Filename)

			// Save the file
			err := ctx.SaveFile(file, fileName)

			if err != nil {
				fmt.Println("Error:", err)
				return err
			}

			// Append the filename to the list
			fileNames = append(fileNames, fileName)
		}
	}

	// Save filenames in Fiber context
	ctx.Locals("fileNames", fileNames)

	return ctx.Next()
}
