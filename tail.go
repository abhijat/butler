package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	progressiveLogEndpoint = "http://localhost:8080/job/%s/%s/logText/progressiveText?start=%d"
)

type jobLog struct {
	log         string
	hasMoreData bool
	offset      int
}

func progressiveLogURL(jobName string, jobNumber string, offset int) string {
	return fmt.Sprintf(progressiveLogEndpoint, jobName, jobNumber, offset)
}

func fetchJobLog(url string) *jobLog {
	response, err := getResponseFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	hasMoreData := len(response.Header.Get("X-More-Data")) != 0
	responseSize := response.Header.Get("X-Text-Size")

	offset, err := strconv.Atoi(responseSize)
	if err != nil {
		log.Fatal(err)
	}

	bytebuf := asBytes(response)
	return &jobLog{string(bytebuf), hasMoreData, offset}
}

func tailJobLatest(jobName string) {
	tailJob(jobName, "lastBuild")
}

func tailJob(jobName string, jobNumber string) {

	offset := 0
	url := progressiveLogURL(jobName, jobNumber, offset)

	for {
		jobData := fetchJobLog(url)
		fmt.Printf(jobData.log)

		if !jobData.hasMoreData {
			break
		} else {
			time.Sleep(200 * time.Millisecond)
			url = progressiveLogURL(jobName, jobNumber, jobData.offset)
		}
	}
}
