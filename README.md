# HTTP Server From Scratch (Go)

[![Go 1.24](https://img.shields.io/badge/go-1.24-blue)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

This repository demonstrates a minimal HTTP/1.1 server built **from scratch** in Go—no `net/http` or third‑party frameworks.
It parses raw TCP streams, implements request‐line and header parsing, gzip encoding, keep‑alive, and basic routing.

[![progress-banner](https://backend.codecrafters.io/progress/http-server/9c0caef2-f580-4de2-8c51-ec862f0811ed)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)


## Features

- Raw TCP listener on port `4221`
- Manual HTTP/1.1 request parser
- Gzip compression when `Accept-Encoding: gzip` is present
- Keep‑alive / `Connection: close` support
- Hand‑written routing for:
  - `GET /` → 200 OK
  - `GET /echo/{msg}` → echoes `{msg}`
  - `GET /user-agent` → returns client’s `User-Agent`
  - `GET|POST /files/{name}` → read/write files from a given directory

## Prerequisites

- Go 1.24+
- GNU make (optional)
- Docker (optional)

## Build & Run

Using Go directly:

```bash
git clone <repo>
cd codecrafters-http-server-go
make build
./bin/server -directory static -port 4221
```

With Docker:

```bash
make docker-build
make docker-run
```

Now open your browser or `curl`:

```bash
curl http://localhost:4221/echo/hello
```

Enjoy your self‑made HTTP server!
