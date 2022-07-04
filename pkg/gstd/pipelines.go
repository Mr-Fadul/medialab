package gstd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/alpheres/medialab/pkg/transfer"
	"github.com/alpheres/medialab/pkg/utils"
	"github.com/rs/zerolog/log"
)

const INPUT_PIPELINE_NAME string = "input_pipeline"
const OUTPUT_PIPELINE_NAME string = "output_pipeline"
const INPUT_PIPELINE_LAUNCH string = "dvbsrc name=dvbsrc ! queue ! interpipesink name=input_intp"
const OUTPUT_PIPELINE_LAUNCH string = "interpipesrc name=output_intp listen-to=input_intp is-live=true ! queue ! tsdemux name=demux ! h264parse ! vaapih264dec ! queue ! videoscale ! capsfilter name=capsfilter ! x264enc name=videoenc ! queue ! mp4mux name=mux ! filesink name=filesink  demux. ! aacparse ! avdec_aac_latm ! queue ! audioconvert ! audioresample ! avenc_aac ! queue ! mux."
const BUS_TIMEOUT string = "1000000000"

func startPipelineInput() (err error) {

	log.Info().Msg("Starting input pipeline")

	setElement(INPUT_PIPELINE_NAME, "dvbsrc", "frequency", os.Getenv("FREQUENCY"))
	setElement(INPUT_PIPELINE_NAME, "dvbsrc", "inversion", os.Getenv("INVERSION"))

	state(INPUT_PIPELINE_NAME, "playing")

	return err
}

func startPipelineOutput() (err error) {

	log.Info().Msg("Starting output pipeline")

	setElement(OUTPUT_PIPELINE_NAME, "capsfilter", "caps", setCapfilter())
	setElement(OUTPUT_PIPELINE_NAME, "videoenc", "bitrate", os.Getenv("VIDEO_BITRATE"))
	setElement(OUTPUT_PIPELINE_NAME, "filesink", "location", utils.GetFileName())
	busFilter(OUTPUT_PIPELINE_NAME, "eos")

	state(OUTPUT_PIPELINE_NAME, "playing")

	return
}

func stopPipelineInput() (err error) {

	log.Info().Msg("Stopping input pipeline")

	state(INPUT_PIPELINE_NAME, "null")

	return
}

func stopPipelineOutput() (err error) {

	log.Info().Msg("Stopping output pipeline")

	busTimeout(OUTPUT_PIPELINE_NAME)

	sendEOS(OUTPUT_PIPELINE_NAME)

	// wait eos
	busReady(OUTPUT_PIPELINE_NAME)

	fName, _ := getElement(OUTPUT_PIPELINE_NAME, "filesink", "location")

	transfer.MoveFileExtension(fName, ".MP4")

	state(OUTPUT_PIPELINE_NAME, "null")

	return
}

func createPipeline(pipeline, launchPipe string) (err error) {

	pipeString := url.QueryEscape(launchPipe)

	url := fmt.Sprintf("pipelines?name=%s&description=%s", pipeline, pipeString)

	method := "POST"

	GSTDNewRequest(url, method)

	return err

}

func delete(pipeline string) (err error) {

	url := fmt.Sprintf("pipelines?name=%s", pipeline)
	method := "DELETE"

	GSTDNewRequest(url, method)

	return err
}

func cleanAllPipelines() {

	log.Info().Msg("Cleaning all pipelines")
	stopPipelineOutput()
	delete(OUTPUT_PIPELINE_NAME)
	stopPipelineInput()
	delete(INPUT_PIPELINE_NAME)
}

func sendEOS(pipeline string) (err error) {

	url := fmt.Sprintf("pipelines/%s/event?name=eos", pipeline)
	method := "POST"

	GSTDNewRequest(url, method)

	return err
}

func setElement(pipeline, element, properties, value string) (err error) {

	url := fmt.Sprintf("pipelines/%s/elements/%s/properties/%s?name=%s", pipeline, element, properties, value)
	method := "PUT"

	GSTDNewRequest(url, method)

	return err
}

func getElement(pipeline, element, properties string) (value string, err error) {

	url := fmt.Sprintf("pipelines/%s/elements/%s/properties/%s", pipeline, element, properties)
	method := "GET"

	body, _ := GSTDNewRequest(url, method)

	ep := ElementProperty{}

	err = json.Unmarshal(body, &ep)
	if err != nil {
		log.Fatal().Err(err).Msg("Error unmarshalling")
	}

	return ep.Response.Value, err
}

func state(pipeline, state string) {

	url := fmt.Sprintf("pipelines/%s/state?name=%s", pipeline, state)
	method := "PUT"

	GSTDNewRequest(url, method)
}

func setCapfilter() (caps string) {

	caps = fmt.Sprintf("video/x-raw,width=%s,height=%s", os.Getenv("WIDTH_VIDEO"), os.Getenv("HEIGHT_VIDEO"))

	return
}

func busTimeout(pipeline string) {

	url := fmt.Sprintf("pipelines/%s/bus/timeout?name=%s", pipeline, BUS_TIMEOUT)
	method := "PUT"

	GSTDNewRequest(url, method)
}

func busReady(pipeline string) {

	url := fmt.Sprintf("pipelines/%s/bus/message", pipeline)
	method := "GET"

	GSTDNewRequest(url, method)
}

func busFilter(pipeline, filter string) {

	url := fmt.Sprintf("pipelines/%s/bus/types?name=%s", pipeline, filter)
	method := "PUT"

	GSTDNewRequest(url, method)
}
