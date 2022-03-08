FROM golang:alpine


WORKDIR /go/src

COPY . .

RUN go mod tidy

EXPOSE $PORT

RUN go build -o app .

ENTRYPOINT ["./app"]
