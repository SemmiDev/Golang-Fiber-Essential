package routes

import (
	"github.com/SemmiDev/fiber-student/controllers"
	"github.com/gofiber/fiber/v2"
)

func StudentRoute(route fiber.Router) {
	route.Get("", controllers.GetStudents)
	route.Post("", controllers.CreateStudent)
	route.Put("/:id", controllers.UpdateStudent)
	route.Delete("/:id", controllers.DeleteStudent)
	route.Get("/:id", controllers.GetStudent)
}