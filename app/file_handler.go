package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type FileHandler struct{ Dir string }

func NewFileHandler(dir string) *FileHandler { return &FileHandler{Dir: dir} }

func (h *FileHandler) Match(path string) bool {
	return strings.HasPrefix(path, "/files/")
}

func (h *FileHandler) Handle(path string, headers map[string]string, body []byte) string {
	name := path[len("/files/"):]
	fp := filepath.Join(h.Dir, name)

	if headers["X-Method"] == "POST" { // we’ll set X-Method in main
		if err := os.WriteFile(fp, body, 0644); err != nil {
			return makeTextResponse("500 Internal Server Error", "")
		}
		return "HTTP/1.1 201 Created\r\n\r\n"
	}

	data, err := os.ReadFile(fp)
	if err != nil {
		return makeTextResponse("404 Not Found", "")
	}
	hdr := "Content-Type: application/octet-stream\r\n" +
		"Content-Length: " + strconv.Itoa(len(data))
	return "HTTP/1.1 200 OK\r\n" + hdr + "\r\n\r\n" + string(data)
}
