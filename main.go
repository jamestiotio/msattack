package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"msattack/config"
	"msattack/utils"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFieldFormat})

	var configuration config.Configuration

	viper.SetConfigFile("./config/config.yml")

	if err := viper.ReadInConfig(); err != nil {
		var configNotFoundError viper.ConfigFileNotFoundError
		if ok := errors.As(err, &configNotFoundError); ok {
			log.Fatal().Err(err).Msg("Error reading config file. Please create a config.yml file in the config folder.")
		} else {
			log.Fatal().Err(err).Msg("Error reading config file. Please ensure that the config.yml file is readable.")
		}
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatal().Err(err).Msg("Error decoding config file. Please check that the config.yml file is valid and properly formatted as a YAML file.")
	}

	if configuration.Port <= 0 {
		log.Fatal().Msg("Port in config file must be a positive integer.")
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

	log.Info().Msg("Spinning up server on port " + fmt.Sprintf("%d", configuration.Port) + "...")

	port := fmt.Sprintf(":%d", configuration.Port)

	app.Get("/", utils.WelcomeMessage)

	err := app.Listen(port)
	if err != nil {
		log.Fatal().Err(err).Msg("Error in starting up server!")
	}
}
