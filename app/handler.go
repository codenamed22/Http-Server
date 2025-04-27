package main

// Handler knows how to Match & Handle a path.
type Handler interface {
	Match(path string) bool
	Handle(path string, headers map[string]string, body []byte) string
}

var routes = map[string][]Handler{}

// Register attaches a path‐handler under an HTTP method.
func Register(method string, h Handler) {
	routes[method] = append(routes[method], h)
}

// RouteRequest looks up handlers for this method, then matches path.
func RouteRequest(method, path string, headers map[string]string, body []byte) string {
	for _, h := range routes[method] {
		if h.Match(path) {
			return h.Handle(path, headers, body)
		}
	}
	return makeTextResponse("404 Not Found", "")
}
