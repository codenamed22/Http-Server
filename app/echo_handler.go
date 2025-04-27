package main

import (
	"strings"
)

type EchoHandler struct{}

func (h *EchoHandler) Match(path string) bool {
	return strings.HasPrefix(path, "/echo/")
}

func (h *EchoHandler) Handle(path string, headers map[string]string, body []byte) string {
	msg := path[len("/echo/"):]
	// always return plain-text; let main.go do gzip encoding if needed
	return makeTextResponse("200 OK", msg)
}

func init() {
	Register("GET", &EchoHandler{})
}
