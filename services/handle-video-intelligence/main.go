package main

import (
	"encoding/json"
	"io/ioutil"

	"net/http"
	"os"

	"github.com/alpheres/medialab/pkg/gcp"
	"github.com/alpheres/medialab/pkg/utils"
	"github.com/rs/zerolog/log"
)

func main() {

	utils.ConfigLog()

	log.Info().Msgf("Starting handle-video-intelligence, Version: %s", os.Getenv("APP_VERSION"))

	http.HandleFunc("/", GetGCSEvent)
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	// Start HTTP server.
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

// GetGCSEvent handles the GCS event.
func GetGCSEvent(w http.ResponseWriter, r *http.Request) {
	var g gcp.GCSEvent
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Info().Msgf("ioutil.ReadAll: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &g); err != nil {
		log.Info().Msgf("json.Unmarshal: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = gcp.HandleVideoIntelligence(r.Context(), g)
	if err != nil {
		log.Info().Msgf("pkg.HandleVideoIntelligence: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}
