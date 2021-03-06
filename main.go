package main

import (
	"github.com/SemmiDev/fiber-student/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main(){
	app := fiber.New()
	app.Use(logger.New())

	routes.StudentRoute(app)

	err := app.Listen(":9090")

	if err != nil {
		panic(err)
	}
}