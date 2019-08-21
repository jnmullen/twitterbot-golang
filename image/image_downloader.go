package image

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// DownloadImage into a byte array from a given URL
func DownloadImage(URL string) []byte {

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body
}
