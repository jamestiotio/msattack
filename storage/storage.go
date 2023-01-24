package storage

import (
	"errors"
	"fmt"
	"os"

	"msattack/config"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetPackFile(c *fiber.Ctx) error {
	configuration := config.GlobalConfig
	log.Info().Msg(fmt.Sprintf("%s/pack/%d/%s", configuration.DataStorageEndpoint, configuration.PackVersion, c.Params("pack_file_name")))
	filePath := fmt.Sprintf("data/pack/%d/%s", configuration.PackVersion, c.Params("pack_file_name"))
	if _, err := os.Stat(filePath); err == nil {
		return c.SendFile(filePath, true)
	} else if errors.Is(err, os.ErrNotExist) {
		log.Warn().Err(err).Msg("Pack file not found.")
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		// Schrodinger's file! File may or may not exist.
		log.Error().Err(err).Msg("Error in getting pack file.")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
}

func GetDataFile(c *fiber.Ctx) error {
	configuration := config.GlobalConfig
	versionNumber, err := c.ParamsInt("version")
	if err != nil {
		log.Error().Err(err).Msg("Error in getting version number.")
		return c.SendStatus(fiber.StatusBadRequest)
	}
	log.Info().Msg(fmt.Sprintf("%s/%d/%s", configuration.DataStorageEndpoint, versionNumber, c.Params("file_name")))
	filePath := fmt.Sprintf("data/%d/%s", versionNumber, c.Params("file_name"))
	if _, err := os.Stat(filePath); err == nil {
		return c.SendFile(filePath, true)
	} else if errors.Is(err, os.ErrNotExist) {
		log.Warn().Err(err).Msg("Data file not found.")
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		// Schrodinger's file! File may or may not exist.
		log.Error().Err(err).Msg("Error in getting data file.")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
}
