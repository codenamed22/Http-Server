package main

type RootHandler struct{}

func (h *RootHandler) Match(path string) bool {
	return path == "/"
}

func (h *RootHandler) Handle(path string, headers map[string]string, body []byte) string {
	return "HTTP/1.1 200 OK\r\n\r\n"
}

func init() {
	Register("GET", &RootHandler{})
}
