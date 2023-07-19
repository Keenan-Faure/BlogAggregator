package fetch

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"objects"
	"utils"
)

func FetchFeed(url string) (objects.RSS, error) {
	if !utils.CheckStringWithWord(url, ".xml") {
		return objects.RSS{}, errors.New("unable to parse non-xml feed")
	}
	resp, err := http.Get(url)
	if err != nil {
		return objects.RSS{}, err
	}
	defer resp.Body.Close()
	body, err_ := io.ReadAll(resp.Body)
	if err_ != nil {
		return objects.RSS{}, err
	}
	result := objects.RSS{}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		return objects.RSS{}, err
	}
	return result, nil
}
