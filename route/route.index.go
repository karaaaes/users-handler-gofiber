package route

import (
	"go-fiber/handler"
	"go-fiber/middleware"
	"go-fiber/model/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	// Assets
	r.Static("/public", "./public/asset")

	// Login
	r.Post("/api/login", handler.AuthLogin)

	r.Get("/api/user", middleware.UserMiddleware, handler.GetAllUser)
	r.Get("/api/user/:id", handler.GetUserById)
	r.Post("/api/user", handler.CreateUser)
	r.Put("/api/user/:id", handler.UpdateUserById)
	r.Put("/api/user/:id/change-password", handler.UpdatePasswordById)
	r.Delete("/api/user/:id", handler.DeleteUser)

	// Upload Comic
	r.Post("api/comic", utils.HandleCoverComic, handler.CreateComic)
}
