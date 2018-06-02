package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func encodeUsernameAndPassword(username, password string) string {
	s := fmt.Sprintf("%s:%s", username, password)
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func addAuthHeader(request *http.Request, username string, password string) {
	authHeader := fmt.Sprintf("Basic %s", encodeUsernameAndPassword(username, password))
	request.Header.Add("Authorization", authHeader)
}

func acquireCredentials() (string, string) {
	return os.Getenv("BUTLER_USER"), os.Getenv("BUTLER_PASSWORD")
}

func acquireCrumb() (string, string) {
	url := `http://localhost:8080/crumbIssuer/api/json`
	crumb, err := getURLContentAsString(url)
	if err != nil {
		log.Fatal("failed to acquire crumb from jenkins: ", err)
	}

	var response struct {
		CrumbRequestField string
		Crumb             string
	}

	err = json.NewDecoder(strings.NewReader(crumb)).Decode(&response)
	if err != nil {
		log.Fatalf("failed to parse crumb response: %v\n", err)
	}

	return response.CrumbRequestField, response.Crumb
}

func addCrumb(request *http.Request) {
	crumbHeader, crumbValue := acquireCrumb()
	request.Header.Set(crumbHeader, crumbValue)
}
