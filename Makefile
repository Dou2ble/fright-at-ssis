# Variables
BINARY_NAME=game

# Default build for the current platform
all: build

# Build for the current platform
build:
	go build -o $(BINARY_NAME) $(GOFILE)

# Build for Windows
windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe .

# Build for Linux
linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) .

# Clean up
clean:
	rm -fv $(BINARY_NAME) $(BINARY_NAME).exe

.PHONY: all build windowsn linux clean