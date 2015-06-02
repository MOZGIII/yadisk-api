package yadisk

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// PerformUpload does the actual upload via unscoped PUT request.
func PerformUpload(url string, data io.Reader) error {
	req, err := http.NewRequest("PUT", url, data)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("upload error [%d]: %s", resp.StatusCode, string(body[:]))
	}
	return nil
}
