FROM golang:1.18-alpine as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o main

FROM alpine:3.11.3

COPY --from=builder /app/main /main

ENTRYPOINT ["./main"]