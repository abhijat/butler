package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type jobInfo struct {
	Name string
	URL  string
}

func newJobInfo(name, url string) *jobInfo {
	return &jobInfo{name, url}
}

func getAllJobs() string {
	url := "http://localhost:8080/api/json?tree=jobs[name,url]"
	content, err := getURLContentAsString(url)
	if err != nil {
		log.Fatal(err)
	}

	var response struct {
		Jobs []jobInfo
	}

	err = json.NewDecoder(strings.NewReader(content)).Decode(&response)
	if err != nil {
		log.Fatal("failed to convert jenkins response to job info: ", err)
	}

	for _, job := range response.Jobs {
		b, _ := json.MarshalIndent(job, "", "  ")
		fmt.Println(string(b))
	}

	return ""
}
