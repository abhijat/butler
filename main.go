package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	jobName := "axon-ca"
	status := getJobStatus(jobName, "lastBuild")
	b, err := json.MarshalIndent(status, "", "  ")
	if err != nil {
		log.Fatal("failed to marshal response:", err)
	}

	fmt.Println(string(b))
}
