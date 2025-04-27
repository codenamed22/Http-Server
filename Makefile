BINARY := bin/server
APP_DIR := app
STATIC := static

.PHONY: all build run clean docker-build docker-run

all: build

build:
	@mkdir -p $(dir $(BINARY))
	go build -o $(BINARY) $(APP_DIR)

run: build
	./$(BINARY) -directory=$(STATIC) -port=4221

clean:
	rm -rf bin/ tmp/

docker-build:
	docker build -t http-server-go .

docker-run:
	docker run --rm -p 4221:4221 -v "$$(pwd)/static:/app/static" http-server-go
