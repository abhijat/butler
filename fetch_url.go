package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func asBytes(response *http.Response) []byte {
	var buf bytes.Buffer
	io.Copy(&buf, response.Body)
	return buf.Bytes()
}

func prettifyResponse(source []byte) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, source, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(buf.Bytes())
}

func getResponseFromURL(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	username, password := acquireCredentials()
	addAuthHeader(request, username, password)

	client := http.Client{}
	response, err := client.Do(request)

	if !strings.HasPrefix(response.Status, "2") {
		errorDescription, _ := ioutil.ReadAll(response.Body)
		err := newHTTPError(response.StatusCode, string(errorDescription))
		return nil, err
	}

	return response, err
}

func getURLContentAsString(url string) (string, error) {
	response, err := getResponseFromURL(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data), nil
}
