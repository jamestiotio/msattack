package utils

import (
	"time"

	"msattack/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

var CUSTOM_TIME_FORMAT string = "2006-01-02 15:04:05"

type ErrorCode struct {
	ErrorCode errors.Error `json:"error_code"`
}

type ErrorCodeWithTime struct {
	ErrorCode errors.Error `json:"error_code"`
	// For all intents and purposes, now_time and server_time are the same.
	NowTime    string `json:"now_time"`
	ServerTime string `json:"server_time"`
}

func WelcomeMessage(c *fiber.Ctx) error {
	err := c.SendString("Welcome to Metal Slug Attack private server! This server was developed by James Raphael Tiovalen. Enjoy and have fun!")
	if err != nil {
		log.Fatal().Err(err).Msg("Error in welcome message. Please check that the server is running properly.")
	}
	return err
}

func DefaultErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).SendString(err.Error())
}

func GenerateErrorCode(status errors.Error) ErrorCode {
	return ErrorCode{
		ErrorCode: status,
	}
}

func GenerateErrorCodeWithTime(status errors.Error) ErrorCodeWithTime {
	return ErrorCodeWithTime{
		ErrorCode:  status,
		NowTime:    time.Now().Format(CUSTOM_TIME_FORMAT),
		ServerTime: time.Now().Format(CUSTOM_TIME_FORMAT),
	}
}
