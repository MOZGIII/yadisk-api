package yadisk

import (
	"io"
	"net/http"
)

// RootAddr is the base URL to Yandex Disk API.
const RootAddr = "https://cloud-api.yandex.net"

// API object incapsulates Yandex Disk REST API access.
type API struct {
	OAuthToken string
	HTTPClient *http.Client
}

// NewAPI constructs a new YaDisk API object.
func NewAPI(oAuthToken string) *API {
	return &API{
		oAuthToken,
		http.DefaultClient,
	}
}

func (a *API) setRequestScope(req *http.Request) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "OAuth "+a.OAuthToken)
}

func (a *API) scopedRequest(method, urlPath string, body io.Reader) (*http.Request, error) {
	fullURL := RootAddr
	if urlPath[:1] != "/" {
		fullURL += "/" + urlPath
	} else {
		fullURL += urlPath
	}

	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return req, err
	}

	a.setRequestScope(req)
	return req, nil
}
