package transfer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/alpheres/medialab/pkg/utils"
	"github.com/rs/zerolog/log"
)

func TransferFiles() {

	transferFiles, exist := os.LookupEnv("TRANSFER_FILES")
	if !exist {
		log.Fatal().Msg("TRANSFER_FILES is not set")
	}

	if transferFiles == "true" {

		transferMethod, exist := os.LookupEnv("TRANSFER_FILE_METHOD")
		if !exist {
			log.Fatal().Msg("TRANSFER_FILE_METHOD is not set")
		}

		switch transferMethod {
		case "azure":
			log.Info().Msg("Transferring files, method: azureFileBlob")
			go SendFileAzureBlobStorage()
		case "gcp":
			log.Info().Msg("Transferring files, method: GCP Bucket")
			go SendFileGCPBucket()
		default:
			log.Fatal().Msgf("Invalid transfer method informed: %s, please inform: azure or gcp", transferMethod)
		}

	} else {
		log.Warn().Msg("Transfer files is disabled")
	}
	// Even if the files are not sent, they will need to be deleted
	go RemoveOldFiles()
}

func GetObjectName(fName string) string {

	channelLocation := os.Getenv("CHANNEL_LOCATION")
	channelName := os.Getenv("CHANNEL_NAME")

	date := time.Now().Format("2006-01-02")

	bname := fmt.Sprintf("%s/%s/%s/%s", channelLocation, channelName, date, path.Base(fName))

	return bname
}

func ReadFile(filePath string) ([]byte, error) {
	dat, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	} else {
		return dat, nil
	}
}

func WalkDir(root string) ([]string, []os.FileInfo, error) {
	var files []string
	var filesInfo []os.FileInfo

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
			filesInfo = append(filesInfo, info)
		}
		return nil
	})

	return files, filesInfo, err
}

func MoveFileExtension(fName string, extension string) (err error) {

	ext := path.Ext(fName)
	outfile := fName[0:len(fName)-len(ext)] + extension

	err = os.Rename(fName, outfile)
	if err != nil {
		log.Fatal().Msgf("Error renaming file: %s", err)
	}

	log.Info().Msgf("Moving file %s", outfile)

	return

}

func RemoveOldFiles() {

	log.Info().Msg("Checking remove old files")

	cutoff := time.Hour * utils.GetFileRemovalTime()

	for {

		_, fileInfo, err := WalkDir(os.Getenv("RECORD_PATH"))
		if err != nil {
			log.Fatal().Msgf("Error has occured: %s", err.Error())
		}

		now := time.Now()
		for _, info := range fileInfo {
			if strings.Contains(info.Name(), ".SENT") {
				if diff := now.Sub(info.ModTime()); diff > cutoff {
					log.Info().Msgf("Deleting %s which is %s old", info.Name(), diff)
					// Delete file if it is older than cutoff
					filefullName := utils.GetAbsoluteFilePath(info.Name())
					err := os.Remove(filefullName)
					if err != nil {
						log.Error().Msgf("Error deleting file: %s", err.Error())
					} else {
						log.Info().Msgf("File: %s deleted successfully", filefullName)
					}
				}
			}
		}

		time.Sleep(time.Second * 30)
	}
}
