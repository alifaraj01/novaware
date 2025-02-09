.PHONY: build install clean

VERSION ?= 0.1.0
COMMIT = $(shell git rev-parse --short HEAD || echo "unknown")
LDFLAGS = -X main.version=${VERSION} -X main.commit=${COMMIT}

build:
	@echo "Building novaware..."
	@go build -v -ldflags "${LDFLAGS}" -o bin/novaware cmd/novaware/main.go

install: build
	@echo "Installing novaware..."
	@sudo cp bin/novaware /usr/local/bin/
	@sudo chmod +x /usr/local/bin/novaware

clean:
	@echo "Cleaning..."
	@rm -rf bin/ 