FROM golang:1.16.7-alpine3.14

WORKDIR /app/

COPY main.go /app/
COPY go.mod /app/

RUN go build

CMD ./http2telegram
