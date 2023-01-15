package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func WelcomeMessage(c *fiber.Ctx) error {
	err := c.SendString("Welcome to Metal Slug Attack private server! This server was developed by James Raphael Tiovalen. Enjoy and have fun!")
	if err != nil {
		log.Fatalf("Error in welcome message: %s", err)
	}
	return err
}

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).SendString(err.Error())
}
