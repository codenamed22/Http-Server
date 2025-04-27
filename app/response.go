package main

import "strconv"

// makeTextResponse builds a plain‐text HTTP/1.1 response with the given status and body.
func makeTextResponse(status, body string) string {
	headers := "Content-Type: text/plain\r\nContent-Length: " + strconv.Itoa(len(body))
	return "HTTP/1.1 " + status + "\r\n" + headers + "\r\n\r\n" + body
}
