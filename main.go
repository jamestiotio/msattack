package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func welcome(c *fiber.Ctx) error {
	err := c.SendString("Welcome to Metal Slug Attack private server!")
	if err != nil {
		log.Fatalf("Error in welcome message: %s", err)
	}
	return err
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	port := ":42069"

	args := os.Args
	if len(args) > 1 {
		port = args[1]
		fmt.Println("Spinning up server on port:", port)
		return
	}

	app.Get("/", welcome)

	err := app.Listen(port)
	if err != nil {
		log.Fatalf("Error in starting up server: %s", err)
	}
}
