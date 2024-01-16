package route

import (
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", "./public/asset")
	r.Get("/api/user", handler.GetAllUser)
	r.Get("/api/user/:id", handler.GetUserById)
	r.Post("/api/user", handler.CreateUser)
	r.Put("/api/user/:id", handler.UpdateUserById)
	r.Put("/api/user/:id/change-password", handler.UpdatePasswordById)
	r.Delete("/api/user/:id", handler.DeleteUser)
}
