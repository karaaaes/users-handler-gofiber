package main

import (
	"go-fiber/database"
	"go-fiber/database/migration"
	"go-fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initial DB
	database.DatabaseInit()

	// Initial Migration
	migration.RunMigration()

	// Initial Route
	route.RouteInit(app)

	app.Listen(":3030")
}
