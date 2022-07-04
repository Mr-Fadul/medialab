package gstd

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alpheres/medialab/pkg/utils"
	"github.com/rs/zerolog/log"
)

func init() {
	go initSignals()
}

func RecController() {

	// Create and start pipelines
	createPipeline(INPUT_PIPELINE_NAME, INPUT_PIPELINE_LAUNCH)
	startPipelineInput()
	createPipeline(OUTPUT_PIPELINE_NAME, OUTPUT_PIPELINE_LAUNCH)
	startPipelineOutput()

	if utils.GetSliceEveryHalfHour() {

		log.Info().Msg("Slice every half hour")

		nextHalfHour := utils.NextHalfHour()

		log.Info().Msgf("The next file will be interrupted in: %s", nextHalfHour)

		for {
			dateNow := time.Now()

			if dateNow.After(nextHalfHour) {

				breakFile()

				nextHalfHour = utils.NextHalfHour()

				log.Info().Msgf("The next file will be interrupted in: %s", nextHalfHour)
			}
			//log.Info().Msgf("Time Now is: %s", timefmt.Format(dateNow, "%Y/%m/%d %H:%M:%S"))
			time.Sleep(time.Second)
		}
	} else {

		log.Info().Msg("Slice every X minutes")

		log.Info().Msgf("The next file will be interrupted in: %s", utils.GetTimeToSliceInfo())

		for {
			time.Sleep(time.Minute * utils.GetTimeToSlice())

			breakFile()

			log.Info().Msgf("The next file will be interrupted in: %s", utils.GetTimeToSliceInfo())
		}
	}

}

func breakFile() {

	log.Info().Msg("Stopping files in the secondary pipeline")

	stopPipelineOutput()
	delete(OUTPUT_PIPELINE_NAME)
	createPipeline(OUTPUT_PIPELINE_NAME, OUTPUT_PIPELINE_LAUNCH)
	startPipelineOutput()
}

// signal handler
func signalHandler(signal os.Signal) {

	log.Logger.Info().Msgf("Received signal: %s", signal)
	log.Logger.Info().Msg("Wait for 1 second to finish processing")
	time.Sleep(time.Second * 1)

	switch signal {

	case syscall.SIGHUP: // kill -SIGHUP XXXX
		log.Logger.Info().Msg("SIGHUP signal received")
		cleanAllPipelines()

	case syscall.SIGINT: // kill -SIGINT XXXX or Ctrl+c
		log.Logger.Info().Msg("SIGINT signal received")
		cleanAllPipelines()

	case syscall.SIGTERM: // kill -SIGTERM XXXX
		log.Logger.Info().Msg("SIGTERM signal received")
		cleanAllPipelines()

	case syscall.SIGQUIT: // kill -SIGQUIT XXXX

		log.Logger.Info().Msg("SIGQUIT signal received")
		cleanAllPipelines()

	default:
		log.Logger.Info().Msg("Unknown signal received")
		cleanAllPipelines()
	}

	log.Logger.Info().Msg("Finished cleanup")
	os.Exit(0)
}

// initialize signal handler
func initSignals() {
	var captureSignal = make(chan os.Signal, 1)
	signal.Notify(captureSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	signalHandler(<-captureSignal)
}
