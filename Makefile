# Target operating system and architecture
GOOS=darwin
GOARCH=arm64
OUTPUT=myBusiness # Output file name
LDFLAGS=-s -w # Go build flags

# Default target
all: build

# Build target
build: main.go
	@echo "Building..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="$(LDFLAGS)" -o $(OUTPUT) main.go

# Clean target
clean:
	@echo "Cleaning..."
	@rm -f $(OUTPUT)

# Run target
run: build
	@echo "Running..."
	@./$(OUTPUT)

.PHONY: all build clean run 