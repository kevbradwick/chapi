package chapi

import (
	"fmt"
	"net/http"
)

func makeRequest(method, path string) (resp *http.Response, err error) {
	client := http.DefaultClient
	url := fmt.Sprintf("%s%s", BaseUrl(), path)
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	request.SetBasicAuth(ApiKey(), "")
	return client.Do(request)
}
