package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type jobStatus struct {
	IsBuilding  bool
	BuildNumber int
	QueueID     int
	Result      string
	URL         string
}

func jobStatusURL(jobName, jobNum string) string {
	return fmt.Sprintf("http://localhost:8080/job/%s/%s/api/json", jobName, jobNum)
}

func getJobStatus(jobName, jobNum string) *jobStatus {

	url := jobStatusURL(jobName, jobNum)
	data, err := getURLContentAsString(url)

	if err != nil {
		log.Fatal("failed to get job state: ", err)
	}

	var status jobStatus
	err = json.NewDecoder(strings.NewReader(data)).Decode(&status)
	if err != nil {
		log.Fatal("error decoding job status response:", err)
	}

	return &status
}
