package logg

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func parseRequestID(ctx ...context.Context) string {
	if len(ctx) != 1 {
		return ""
	}
	requestID := ctx[0].Value("x-telegram-id")
	if requestID == nil {
		return ""
	}
	return requestID.(string)
}

func Info(ctx ...context.Context) *zerolog.Event {
	requestID := parseRequestID(ctx...)
	if requestID != "" {
		return log.Info().Any("x-telegram-id", requestID)
	}
	return log.Info()
}

func Fatal(ctx ...context.Context) *zerolog.Event {
	requestID := parseRequestID(ctx...)
	if requestID != "" {
		return log.Fatal().Any("x-chat-id", requestID)
	}
	return log.Fatal()
}

func Error(ctx ...context.Context) *zerolog.Event {
	requestID := parseRequestID(ctx...)
	if requestID != "" {
		return log.Error().Any("x-chat-id", requestID)
	}
	return log.Error()
}
