package main

import (
	"os"

	"github.com/alpheres/medialab/pkg/gstd"
	"github.com/alpheres/medialab/pkg/transfer"
	"github.com/alpheres/medialab/pkg/utils"
	"github.com/rs/zerolog/log"
)

func main() {

	utils.ConfigLog()

	log.Info().Msgf("Starting digitaltv-recorder, Version: %s", os.Getenv("APP_VERSION"))

	transfer.TransferFiles()

	gstd.RecController()
}
