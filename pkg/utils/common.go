package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/itchyny/timefmt-go"
	"github.com/rs/zerolog/log"
)

// Cria-se um diretorio para gravação caso não exista
func createDirectory(path string) (dirNew string) {

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal().Msgf("Error creating directory:: %s", err)
	}

	return
}

// Responsável em verificar qual a próxima meia hora exata para corte de arquivos
func NextHalfHour() (now time.Time) {

	now = time.Now()

	if now.Minute() < 30 {
		now = now.Truncate(time.Minute * 30).Add(time.Minute * 30)
	} else {
		now = now.Truncate(time.Hour * 1).Add(time.Hour * 1)
	}

	return
}

func GetFileName() (name string) {

	fullPath := fmt.Sprintf("%s/%s", os.Getenv("RECORD_PATH"), os.Getenv("CHANNEL_NAME"))

	createDirectory(fullPath)

	initDate := time.Now()

	var finalDate time.Time

	if GetSliceEveryHalfHour() {
		finalDate = NextHalfHour()
	} else {
		finalDate = time.Now().Add(time.Minute * GetTimeToSlice())
	}

	name = fmt.Sprintf(
		"%s/%s-%s-%s.TEMP",
		fullPath,
		os.Getenv("CHANNEL_NAME"),
		timefmt.Format(initDate, "%Y%m%d_%H%M%S"),
		timefmt.Format(finalDate, "%Y%m%d_%H%M%S"))

	return name
}

func GetTimeToSlice() time.Duration {

	t, err := strconv.ParseInt(os.Getenv("TIME_TO_SLICE"), 10, 64)
	if err != nil {
		log.Fatal().Msgf("Error getting time to slice: %s", err)
	}

	return time.Duration(t)
}

func GetTimeToSliceInfo() string {

	t := time.Now().Add(time.Minute * GetTimeToSlice())

	return timefmt.Format(t, "%Y/%m/%d %H:%M:%S")
}

func GetSliceEveryHalfHour() bool {

	return os.Getenv("SLICE_EVERY_HALF_HOUR") == "true"
}

func GetFileRemovalTime() time.Duration {

	t, err := strconv.ParseInt(os.Getenv("FILE_REMOVAL_TIME"), 10, 64)
	if err != nil {
		log.Fatal().Msgf("Error getting file removal time: %s", err)
	}

	return time.Duration(t)
}

func GetAbsoluteFilePath(file string) string {

	return fmt.Sprintf("%s/%s/%s", os.Getenv("RECORD_PATH"), os.Getenv("CHANNEL_NAME"), file)
}
