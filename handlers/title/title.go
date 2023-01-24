package title

import (
	"fmt"

	"msattack/config"
	"msattack/errors"
	"msattack/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

var PACK_INFO_URL = `https://%s%s/pack/%d/`

func GetPackInfo(c *fiber.Ctx) error {
	log.Info().Msg("POST /title/get_pack_info")

	configuration := config.GlobalConfig

	actualPackInfoURL := fmt.Sprintf(PACK_INFO_URL, configuration.StorageDomain, configuration.DataStorageEndpoint, configuration.PackVersion)

	c.Set("Connection", "close")
	c.Set("Content-Encoding", "gzip")
	c.Set("Server", "Apache")
	c.Set("Vary", "Accept-Encoding")

	err := c.Status(fiber.StatusOK).JSON(fiber.Map{
		"version":  configuration.PackVersion,
		"url":      actualPackInfoURL,
		"response": utils.GenerateErrorCode(errors.SUCCESS),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error in getting pack info.")
	}

	return err
}

func GetFileList(c *fiber.Ctx) error {
	log.Info().Msg("POST /title/get_file_list")

	configuration := config.GlobalConfig

	c.Set("Connection", "close")
	c.Set("Content-Encoding", "gzip")
	c.Set("Server", "Apache")
	c.Set("Vary", "Accept-Encoding")

	err := c.Status(fiber.StatusOK).JSON(fiber.Map{
		"master_ver": configuration.MasterVersion,
		// Not sure what these are for (parallel download limits for master table and DLC?)
		"max_dl_stream_num_for_mtbl": 1,
		"max_dl_stream_num_for_dlc":  4,
		"response":                   utils.GenerateErrorCode(errors.SUCCESS),
		"file_list":                  generateFileList(),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error in getting file list.")
	}

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
