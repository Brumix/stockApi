FROM golang:alpine

ENV GIN_MODE=debug
ENV PORT=8080

WORKDIR /go/src

COPY . .

RUN go mod tidy

EXPOSE $PORT

RUN go build -o app .

ENTRYPOINT ["./app"]
