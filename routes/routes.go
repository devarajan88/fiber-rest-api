package routes

import (
	"fiber-rest-api/controller"
	"fiber-rest-api/service"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/students", controller.GetStudents)
	app.Get("/student/:id", controller.GetStudentById)
	app.Post("/student", controller.AdmitStudent)
	app.Post("/upload", service.UploadFile)

}
