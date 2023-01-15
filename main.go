package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"msattack/config"
	"msattack/utils"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func main() {
	var (
		configuration config.Configuration
	)

	viper.SetConfigFile("./config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Please create a config.yml file in the config folder.")
		} else {
			fmt.Printf("Please ensure that the config.yml file is readable.")
		}
		os.Exit(1)
	}
	if err := viper.Unmarshal(&configuration); err != nil {
		fmt.Printf("Error decoding config file: %s\n", err)
		fmt.Printf("Please check that the config.yml file is valid and properly formatted as a YAML file.")
		os.Exit(1)
	}

	app := fiber.New(fiber.Config{
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
		ErrorHandler: utils.DefaultErrorHandler,
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

	port := fmt.Sprintf(":%d", configuration.Port)

	args := os.Args
	if len(args) > 1 {
		port = args[1]
		fmt.Println("Spinning up server on port:", port)
		return
	}

	app.Get("/", utils.WelcomeMessage)

	err := app.Listen(port)
	if err != nil {
		log.Fatalf("Error in starting up server: %s", err)
	}
}
