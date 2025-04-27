# Stage 1: build
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./app

# Stage 2: final
FROM scratch
COPY --from=builder /app/server /server
EXPOSE 4221
ENTRYPOINT ["/server"]
CMD ["-directory", "static", "-port", "4221"]
