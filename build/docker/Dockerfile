FROM golang:1.15.2-buster AS dev

WORKDIR /go/src/github.com/purusah/zebra

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/build/zebra
