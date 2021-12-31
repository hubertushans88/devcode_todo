package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/hubertushans88/devcode_todo/controllers/activity"
	"github.com/hubertushans88/devcode_todo/controllers/todo"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	app := fiber.New(fiber.Config{
		Prefork:               true,
		DisableStartupMessage: true,
	})

	//app.Use(logger.New())
	app.Use(cache.New())

	app.Get("/activity-groups", activity.ReadAll)
	app.Post("/activity-groups", activity.Create)
	app.Get("/activity-groups/:id", activity.ReadOne)
	app.Patch("/activity-groups/:id", activity.Update)
	app.Delete("/activity-groups/:id", activity.Delete)

	app.Get("/todo-items", todo.ReadAll)
	app.Post("/todo-items", todo.Create)
	app.Get("/todo-items/:id", todo.ReadOne)
	app.Patch("/todo-items/:id", todo.Update)
	app.Delete("/todo-items/:id", todo.Delete)

	//app.Listen(":5000")
	app.Listen(":3030")
}
