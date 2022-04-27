package main

import (
	"go_mongo_rest_api_fiber/configs"
	"go_mongo_rest_api_fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	// })
	//run database
	configs.ConnectDB()
	//routes
	routes.UserRoute(app)
	app.Listen(":6000")
}
