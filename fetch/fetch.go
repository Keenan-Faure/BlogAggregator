package fetch

import (
	"encoding/xml"
	"io"
	"net/http"
)

func FetchFeed(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err_ := io.ReadAll(resp.Body)
	if err_ != nil {
		return err
	}
	xml.Unmarshal(body, &structure)
	return nil
	// result := pokeloc{}
	// err_r := json.Unmarshal(body, &result)
	// if err_r != nil {
	// 	return pokeloc{}, err_r
	// }
}
