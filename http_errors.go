package main

import (
	"fmt"
)

type httpError struct {
	ErrorCode int
	Text      string
}

func newHTTPError(errorCode int, errorText string) *httpError {
	return &httpError{errorCode, errorText}
}

func (e *httpError) Error() string {
	return fmt.Sprintf("error code: %v\ntext: %v", e.ErrorCode, e.Text)
}
