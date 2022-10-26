package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bytedance/sonic"
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
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Date", time.Now().Format(time.RFC1123))
		c.Set("Server", "Apache")
		c.Set("Vary", "Accept-Encoding")
		c.Set("Content-Encoding", "gzip")
		c.Set("Connection", "close")
		c.Set("MSA-Signature", "https://github.com/jamestiotio/msattack")
		return c.Next()
	})

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
