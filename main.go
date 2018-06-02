package main

import (
	"fmt"
	"log"
	"net/url"
)

func runJob() {
	jobName := "axon-ca"

	params := url.Values{}
	params.Set("run", "foo")

	responseCode, err := triggerBuild(jobName, params)
	if err != nil {
		log.Fatal(err)
	}

	if responseCode == 201 {
		fmt.Printf("triggered build of %s, follow the output below:\n\n", jobName)
		tailJobLatest(jobName)
	}
}

func main() {
	runJob()
}
