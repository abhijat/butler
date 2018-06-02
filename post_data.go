package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postWithParams(postURL string, params url.Values) (int, error) {
	request, err := http.NewRequest("POST", postURL, strings.NewReader(params.Encode()))

	username, password := acquireCredentials()
	addAuthHeader(request, username, password)

	addCrumb(request)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return -1, err
	}

	defer response.Body.Close()
	if !strings.HasPrefix(response.Status, "2") {
		b, _ := ioutil.ReadAll(response.Body)
		return -1, newHTTPError(response.StatusCode, string(b))
	}

	return response.StatusCode, nil
}
