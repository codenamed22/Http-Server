package main

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

// parseRequest reads the request line, headers, then body (if any).
func parseRequest(reader *bufio.Reader) (method, path, proto string, headers map[string]string, body []byte, err error) {
	// Read request line
	line, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.TrimRight(line, "\r\n")
	parts := strings.SplitN(line, " ", 3)
	if len(parts) != 3 {
		err = errors.New("malformed request line")
		return
	}
	method, path, proto = parts[0], parts[1], parts[2]

	// Read headers
	headers = make(map[string]string)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil {
			err = err2
			return
		}
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			break
		}
		kv := strings.SplitN(line, ":", 2)
		if len(kv) != 2 {
			continue
		}
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])
		headers[key] = value
	}
	// Read body if Content-Length present
	if cl, ok := headers["Content-Length"]; ok {
		n, err2 := strconv.Atoi(cl)
		if err2 != nil {
			err = err2
			return
		}
		if n > 0 {
			body = make([]byte, n)
			if _, err2 := io.ReadFull(reader, body); err2 != nil {
				err = err2
				return
			}
		}
	}
	return
}
