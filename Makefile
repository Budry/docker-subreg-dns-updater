OUTPUT=subreg-dns-updater-cli

all: build
build:
		go build -o dist/$(OUTPUT) -v