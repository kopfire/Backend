# Start from a small, secure base image
FROM golang:1.19-alpine

COPY . /go/src/app

WORKDIR /go/src/app/cmd/AvitoBackend

RUN go build -o app main.go

EXPOSE 9999

CMD ["./app"]