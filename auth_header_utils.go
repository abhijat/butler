package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func encodeUsernameAndPassword(username, password string) string {
	s := fmt.Sprintf("%s:%s", username, password)
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func addAuthHeader(request *http.Request, username string, password string) {
	authHeader := fmt.Sprintf("Basic %s", encodeUsernameAndPassword(username, password))
	request.Header.Add("Authorization", authHeader)
}
