package title

import (
	"fmt"

	"msattack/config"
	"msattack/errors"
	"msattack/managers"
	"msattack/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

const PackInfoURL = `https://%s%s/pack/%d/`

func GetPackInfo(c *fiber.Ctx) error {
	log.Info().Msg("POST /title/get_pack_info")

	configuration := config.GlobalConfig

	actualPackInfoURL := fmt.Sprintf(PackInfoURL, configuration.StorageDomain, configuration.DataStorageEndpoint, configuration.PackVersion)

	err := c.Status(fiber.StatusOK).JSON(fiber.Map{
		"version":  configuration.PackVersion,
		"url":      actualPackInfoURL,
		"response": utils.GenerateErrorCode(errors.SUCCESS),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error in getting pack info.")
	}

	c.Set("Connection", "close")
	c.Set("Content-Encoding", "gzip")
	c.Set("Server", "Apache")
	c.Set("Vary", "Accept-Encoding")

	return err
}

func GetFileList(c *fiber.Ctx) error {
	log.Info().Msg("POST /title/get_file_list")

	configuration := config.GlobalConfig

	err := c.Status(fiber.StatusOK).JSON(fiber.Map{
		"master_ver": configuration.MasterVersion,
		// Not sure what these are for (parallel download limits for master table and DLC?)
		"max_dl_stream_num_for_mtbl": 1,
		"max_dl_stream_num_for_dlc":  4,
		"response":                   utils.GenerateErrorCodeWithTime(errors.SUCCESS),
		"file_list":                  managers.GenerateFileList(),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error in getting file list.")
	}

	c.Set("Connection", "close")
	c.Set("Content-Encoding", "gzip")
	c.Set("Server", "Apache")
	c.Set("Vary", "Accept-Encoding")

	return err
}

func GetMasterTable(c *fiber.Ctx) error {
	// Since we can potentially have multiple "table[]" query parameters, we iterate over them to get the full list
	tableNames := c.Context().QueryArgs().PeekMulti("table[]")
	for _, tableName := range tableNames {
		log.Info().Msgf("GET /title/get_master_table?table[]=%s", tableName)
	}
	err := c.Status(fiber.StatusOK).SendString("Coming soon!")
	if err != nil {
		log.Fatal().Err(err).Msg("Error in getting master table data.")
	}
	return err
}
