package transfer

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/rs/zerolog/log"
)

func SendFileAzureBlobStorage() {

	for {

		files, _, err := WalkDir(os.Getenv("RECORD_PATH"))

		if err != nil {
			log.Fatal().Msgf("Error has occured: %s", err.Error())
		} else {

			//if len(files) == 0 {
			//	log.Info().Msg("No files were found to upload to blob storage")
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
				u, errU := uploadBytesToBlob(m[fName], fName)
				if errU != nil {
					log.Error().Msgf("Error during upload: %s ", errU.Error())
				}
				log.Info().Msg("Finished uploading: " + u)

				// Move file extension to sent
				MoveFileExtension(fName, ".SENT")
			}

		}

		// Wait for next iteration
		time.Sleep(time.Second * 5)
	}

}

func GetAccountInfo() (string, string, string, string) {
	azrKey := os.Getenv("STORAGE_ACCOUNT_KEY")
	azrBlobAccountName := os.Getenv("STORAGE_ACCOUNT_NAME")
	azrPrimaryBlobServiceEndpoint := fmt.Sprintf("https://%s.blob.core.windows.net/", azrBlobAccountName)
	azrBlobContainer := os.Getenv("BLOB_STORAGE_NAME")

	return azrKey, azrBlobAccountName, azrPrimaryBlobServiceEndpoint, azrBlobContainer
}

func uploadBytesToBlob(b []byte, fName string) (string, error) {
	azrKey, accountName, endPoint, container := GetAccountInfo()
	u, _ := url.Parse(fmt.Sprint(endPoint, container, "/", GetObjectName(fName)))
	credential, errC := azblob.NewSharedKeyCredential(accountName, azrKey)
	if errC != nil {
		return "", errC
	}

	blockBlobUrl := azblob.NewBlockBlobURL(*u, azblob.NewPipeline(credential, azblob.PipelineOptions{
		Retry: azblob.RetryOptions{
			TryTimeout: time.Second * 30000,
		},
	}))

	ctx := context.Background()
	o := azblob.UploadToBlockBlobOptions{
		BlobHTTPHeaders: azblob.BlobHTTPHeaders{
			ContentType:     "video/mp4",
			ContentLanguage: "pt-BR",
		},
	}

	_, errU := azblob.UploadBufferToBlockBlob(ctx, b, blockBlobUrl, o)
	return blockBlobUrl.String(), errU
}
