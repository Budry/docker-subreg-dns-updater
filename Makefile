OUTPUT=docker-subreg-dns-updater

all: build
build:
		docker build -t budry/docker-subreg-dns-updater .
		docker push budry/docker-subreg-dns-updater