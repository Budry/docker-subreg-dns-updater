FROM golang:alpine as builder

RUN apk add docker

WORKDIR /go/src/app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go install -v ./...


FROM alpine

RUN apk add docker

COPY --from=builder /go/bin/docker-subreg-dns-updater /usr/bin/docker-subreg-dns-updater
RUN ls -la /usr/bin

ENTRYPOINT ["docker-subreg-dns-updater"]