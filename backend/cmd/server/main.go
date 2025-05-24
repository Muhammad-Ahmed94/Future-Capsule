package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/Muhammad-Ahmed94/Future-Capsule/cmd/utils"
)

func main() {
	utils.LoadEnv()              // Load environment variables
	port := utils.GetEnv("PORT") // Gets Env specific key name

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Future Capsule API (Fiber) running")
	})

	log.Printf("Server running port %s...", port)

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
