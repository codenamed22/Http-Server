package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	var directory string
	flag.StringVar(&directory, "directory", "", "files directory")
	flag.Parse()

	// register FileHandler under GET & POST
	fh := NewFileHandler(directory)
	Register("GET", fh)
	Register("POST", fh)

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		go func(c net.Conn) {
			defer c.Close()
			reader := bufio.NewReader(c)
			for {
				// read one request
				method, path, _, hdrs, body, err := parseRequest(reader)
				if err != nil {
					// EOF or parse error → close connection
					return
				}
				// stash method into headers
				hdrs["X-Method"] = method

				// dispatch
				resp := RouteRequest(method, path, hdrs, body)

				// gzip logic (unchanged)
				if strings.Contains(hdrs["Accept-Encoding"], "gzip") && !strings.Contains(resp, "Content-Encoding:") {
					parts := strings.SplitN(resp, "\r\n\r\n", 2)
					hdrLines := strings.Split(parts[0], "\r\n")
					bodyBytes := []byte(parts[1])

					var buf bytes.Buffer
					gz := gzip.NewWriter(&buf)
					gz.Write(bodyBytes)
					gz.Close()
					compressed := buf.Bytes()

					newHdr := []string{hdrLines[0]}
					for _, line := range hdrLines[1:] {
						if strings.HasPrefix(line, "Content-Length:") || strings.HasPrefix(line, "Content-Encoding:") {
							continue
						}
						newHdr = append(newHdr, line)
					}
					newHdr = append(newHdr, "Content-Encoding: gzip")
					newHdr = append(newHdr, "Content-Length: "+strconv.Itoa(len(compressed)))

					resp = strings.Join(newHdr, "\r\n") + "\r\n\r\n" + string(compressed)
				}

				// explicit Connection: close support
				if strings.ToLower(hdrs["Connection"]) == "close" {
					parts := strings.SplitN(resp, "\r\n\r\n", 2)
					resp = parts[0] + "\r\nConnection: close\r\n\r\n" + parts[1]
				}

				// send response
				if _, err := c.Write([]byte(resp)); err != nil {
					return
				}

				// close connection after responding if requested
				if strings.ToLower(hdrs["Connection"]) == "close" {
					return
				}
			}
		}(conn)
	}
}
