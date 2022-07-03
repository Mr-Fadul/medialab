package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func ConfigLogger() {

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("No .env file found")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		log.Fatal().Msg("LOG_LEVEL is not set")
	}

	switch logLevel {

	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Fatal().Msgf("Invalid log level informed: %s, please inform: debug, info, warn, error, fatal or panic", logLevel)
	}

}

func Greet(audience string) string {
	return fmt.Sprintf("Hello, %s!", audience)
}
