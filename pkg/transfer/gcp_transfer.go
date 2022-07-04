package transfer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/rs/zerolog/log"
)

func SendFileGCPBucket() {

	for {

		files, _, err := WalkDir(os.Getenv("RECORD_PATH"))

		if err != nil {
			log.Fatal().Msgf("Error has occured: %s", err.Error())
		} else {

			//if len(files) == 0 {
			//	log.Info().Msg("No files were found to upload to GCP Bucket")
			//}

			var fFiles []string
			for _, fName := range files {
				if strings.Contains(fName, ".MP4") {
					fFiles = append(fFiles, fName)
				}
			}

			m := make(map[string][]byte)

			// Read file contents into memory
			for _, fName := range fFiles {
				log.Info().Msgf("Found file: %s", fName)
				dat, errR := ReadFile(fName)

				if errR != nil {
					log.Error().Msgf("Error reading file: %s Error: ", errR.Error())
				} else {
					log.Info().Msgf("File: %s read successfully", fName)
					m[fName] = dat
				}
			}

			// push file contents from memory to Azure
			for _, fName := range fFiles {
				log.Info().Msgf("Started uploading: %s", fName)
				err := streamFileUpload(m[fName], fName)
				if err != nil {
					log.Error().Msgf("Error during upload: %s ", err.Error())
				}

				// Move file extension to sent
				MoveFileExtension(fName, ".SENT")
			}

		}

		// Wait for next iteration
		time.Sleep(time.Second * 5)
	}
}

// streamFileUpload uploads an object via a stream.
func streamFileUpload(b []byte, object string) error {

	bucket, exist := os.LookupEnv("GCP_BUCKET_NAME")
	if !exist {
		log.Fatal().Msg("GCP_BUCKET_NAME environment variable is not set")
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	buf := bytes.NewBuffer(b)

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(GetObjectName(object)).NewWriter(ctx)
	//wc.ChunkSize = 0 // note retries are not supported for chunk size 0.

	if _, err = io.Copy(wc, buf); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	// Data can continue to be added to the file until the writer is closed.
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	log.Info().Msgf("Finished uploading:  %s", object)

	return nil
}
