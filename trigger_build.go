package main

import (
	"fmt"
	"log"
	"net/url"
)

func triggerBuild(jobName string, params url.Values) (int, error) {

	var url string
	if params == nil {
		url = fmt.Sprintf("http://localhost:8080/job/%s/build", jobName)
	} else {
		url = fmt.Sprintf("http://localhost:8080/job/%s/buildWithParameters", jobName)
	}

	responseCode, err := postWithParams(url, params)
	if err != nil {
		log.Fatal(err)
	}

	return responseCode, err
}
