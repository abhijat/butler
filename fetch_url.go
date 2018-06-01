package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	addAuthHeader(request, os.Getenv("BUTLER_USER"), os.Getenv("BUTLER_PASSWORD"))

	client := http.Client{}
	return client.Do(request)
}

func getURLContentAsString(url string) string {
	response, err := getResponseFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
