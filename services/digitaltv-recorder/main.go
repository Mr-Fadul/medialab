package main

import (
	"os"

	"github.com/alpheres/medialab/pkg/utils"
	"github.com/rs/zerolog/log"
)

func main() {

	utils.ConfigLogger()

	log.Info().Msgf("Starting digitaltv-recorder, Version: %s", os.Getenv("APP_VERSION"))
}
