FROM golang:1.18-alpine
RUN apk add build-base
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /product_q_srv

ENTRYPOINT ["/product_q_srv"]