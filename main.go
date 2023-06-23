package main

import (
	db "fiber-rest-api/config"
	"fiber-rest-api/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Start connect to DB...")
	db.Connect()

	app := fiber.New()

	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "An API call  is made successfully!!!",
		})
	})

	routes.Setup(app)
	app.Listen(":9000")

	//ser.CsvProcessor()
}
