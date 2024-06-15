FROM golang:1.22

WORKDIR /go/app
RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]
