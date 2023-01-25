package main

import (
	"fmt"
	"os"
	"time"

	"msattack/config"
	"msattack/handlers/title"
	"msattack/middleware"
	"msattack/storage"
	"msattack/utils"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFieldFormat})

	config.LoadConfig("config/config.yml")

	configuration := config.GlobalConfig

	if configuration.Port <= 0 {
		log.Fatal().Msg("Port in config file must be a positive integer.")
	}

	app := fiber.New(fiber.Config{
		JSONEncoder:  sonic.Marshal,
		JSONDecoder:  sonic.Unmarshal,
		ErrorHandler: utils.DefaultErrorHandler,
	})

	app.Use(cors.New())

	// Set default headers
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Date", time.Now().Format(time.RFC1123))
		c.Set("MSA-Signature", "https://github.com/jamestiotio/msattack")
		return c.Next()
	})

	// Define global maintenance mode handler
	app.Use(middleware.CheckForMaintenance)

	// Define all other handlers
	app.Get("/", utils.WelcomeMessage)

	titleGroup := app.Group("/title")
	titleGroup.Post("/get_pack_info", title.GetPackInfo)
	titleGroup.Post("/get_file_list", title.GetFileList)
	titleGroup.Get("/get_master_table", title.GetMasterTable)

	storageGroup := app.Group(configuration.DataStorageEndpoint)
	storageGroup.Get("/:version/:file_name", storage.GetDataFile)
	storagePackGroup := storageGroup.Group(fmt.Sprintf("/pack/%d", configuration.PackVersion))
	storagePackGroup.Get("/:pack_file_name", storage.GetPackFile)

	log.Info().Msg("Spinning up server on port " + fmt.Sprintf("%d", configuration.Port) + "...")

	port := fmt.Sprintf(":%d", configuration.Port)

	if configuration.UseTLS {
		err := app.ListenTLS(port, configuration.TLSCertPath, configuration.TLSKeyPath)
		if err != nil {
			log.Fatal().Err(err).Msg("Error in starting up server!")
		}
	} else {
		err := app.Listen(port)
		if err != nil {
			log.Fatal().Err(err).Msg("Error in starting up server!")
		}
	}
}
