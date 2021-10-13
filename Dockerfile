# Base environment (alias: base)
FROM amd64/golang:1.16.4 AS base

# Development backend environment
FROM base as dev
WORKDIR /root
RUN apt-get update && apt-get install -y fswatch
RUN go get github.com/go-delve/delve/cmd/dlv
WORKDIR /anonym

FROM base as builder
COPY . /anonym
WORKDIR /anonym
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

#Production environment
FROM scratch as prod
COPY --from=builder /anonym/config/config.yml /anonym/config/config.yml
COPY --from=builder /anonym/anonym /anonym/anonym
WORKDIR /anonym

