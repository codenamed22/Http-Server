package main

type UserAgentHandler struct{}

func (h *UserAgentHandler) Match(path string) bool {
	return path == "/user-agent"
}

func (h *UserAgentHandler) Handle(path string, headers map[string]string, body []byte) string {
	return makeTextResponse("200 OK", headers["User-Agent"])
}

func init() {
	Register("GET", &UserAgentHandler{})
}
