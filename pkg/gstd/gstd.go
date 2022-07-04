package gstd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

func GSTDNewRequest(url string, method string) (body []byte, err error) {

	fullUrl := fmt.Sprintf("%s%s", gstdApiUrl(), url)

	client := &http.Client{}
	req, err := http.NewRequest(method, fullUrl, nil)

	if err != nil {
		log.Fatal().Msgf("Error creating request: %s", err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal().Msgf("Error sending request: %s", err)
	}

	defer res.Body.Close()

	log.Debug().
		Str("method", res.Request.Method).
		Str("path", res.Request.URL.String()).
		RawJSON("body", body).
		Int("status", res.StatusCode)

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func gstdApiUrl() (urlApi string) {

	urlApi = fmt.Sprintf("%s/", os.Getenv("GSTD_API_URL"))

	return
}
