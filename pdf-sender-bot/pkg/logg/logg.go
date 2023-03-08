package logg

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type IDs struct {
	requestID  string
	telegramID string
}

func parseIDs(ctx ...context.Context) *IDs {

	ids := &IDs{}

	if len(ctx) != 1 {
		return ids
	}

	telegramID := ctx[0].Value("x-telegram-id")
	if telegramID != nil {
		ids.telegramID = telegramID.(string)
	}

	requestID := ctx[0].Value("x-request-id")
	if requestID != nil {
		ids.requestID = requestID.(string)
	}

	return ids
}

func fillIDs(logger *zerolog.Event, ctx ...context.Context) *zerolog.Event {
	ids := parseIDs(ctx...)

	if ids.telegramID != "" {
		logger = logger.Any("x-telegram-id", ids.telegramID)
	}

	if ids.requestID != "" {
		logger = logger.Any("x-request-id", ids.requestID)
	}
	return logger
}

func Info(ctx ...context.Context) *zerolog.Event {
	return fillIDs(log.Info(), ctx...)
}

func Fatal(ctx ...context.Context) *zerolog.Event {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Stack().Logger()
	return fillIDs(logger.Fatal(), ctx...)
}

func Error(ctx ...context.Context) *zerolog.Event {
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Stack().Logger()
	return fillIDs(logger.Error(), ctx...)
}
