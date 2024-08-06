.PHONY: run build css watch dev

# Go binary name
BINARY_NAME=marked

# Main Go package
MAIN_PACKAGE=./

# Build the project
build:
	go build -o ${BINARY_NAME} ${MAIN_PACKAGE}

# Run the project
run:
	go run ${MAIN_PACKAGE}

# Build CSS
css:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css

# Watch CSS changes
watch-css:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

run-templ:
	templ generate --watch

run-air:
	air
# Run development mode
dev:
	make run-air run-templ run watch-css

# Clean up
clean:
	go clean
	rm ${BINARY_NAME}

# Test the project
test:
	go test ./...

# Generate templ files
generate:
	templ generate

