FROM golang:1.27-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 go build -o /pong .

FROM alpine:3.20
COPY --from=builder /pong /pong
ENTRYPOINT ["/pong"]
