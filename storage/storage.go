package storage

import (
	"fmt"

	"msattack/config"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetPackFile(c *fiber.Ctx) error {
	configuration := config.GlobalConfig
	log.Info().Msg(fmt.Sprintf("%s/pack/%d/%s", configuration.DataStorageEndpoint, configuration.PackVersion, c.Params("pack_file_name")))
	filepath := fmt.Sprintf("data/pack/%d/%s", configuration.PackVersion, c.Params("pack_file_name"))

	err := c.SendFile(filepath, true)

	if err != nil {
		log.Warn().Msg("Error in getting pack file. Ensure that the pack file exists and is accessible.")
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		return err
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
	filepath := fmt.Sprintf("data/%d/%s", versionNumber, c.Params("file_name"))

	err = c.SendFile(filepath, true)

	if err != nil {
		log.Warn().Msg("Error in getting data file. Ensure that the data file exists and is accessible.")
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		return err
	}
}
