package main

import (
	"aprendendo/web_server/src/controller"
	"fmt"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	fmt.Println("Server started...")

	app.Post("user/posts", func(ctx *fiber.Ctx) {
		controller.GetUserHandler(ctx)
	})

	app.Listen(8081)
}
