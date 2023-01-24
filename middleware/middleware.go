package middleware

import (
	"msattack/config"
	"msattack/errors"
	"msattack/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func CheckForMaintenance(c *fiber.Ctx) error {
	configuration := config.GlobalConfig

	if configuration.IsMaintenanceMode {
		err := c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"response": utils.GenerateErrorCode(errors.MAINTENANCE),
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Error in indicating maintenance mode.")
		}
		return err
	} else {
		return c.Next()
	}
}
