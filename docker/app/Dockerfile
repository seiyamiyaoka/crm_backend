FROM golang:1.19-alpine

WORKDIR /go/src
COPY ./src .

RUN apk upgrade --update && \
    apk add --no-cache git alpine-sdk build-base

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

RUN go mod tidy
CMD ["air", "-c", ".air.toml"]