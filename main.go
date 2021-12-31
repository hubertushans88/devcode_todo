package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hubertushans88/devcode_todo/controllers/activity"
)

func main(){
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(logger.New())

	app.Get("/activity-groups", activity.ReadAll)
	app.Post("/activity-groups", activity.Create)
	app.Get("activity-groups/:id", activity.ReadOne)
	app.Patch("activity-groups/:id", activity.Update)
	app.Delete("activity-groups/:id", activity.Delete)

	app.Listen(":5000")
}