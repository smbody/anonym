# Development environment
# Unfortunately, linux alpine doesn't have fswatch package by default, so we will need to download source code and make it by outselves.
FROM golang:alpine AS dev
RUN apk update && apk upgrade && \
apk add --no-cache bash git openssh autoconf automake libtool gettext gettext-dev make g++ texinfo curl
WORKDIR /root
RUN wget https://github.com/emcrisostomo/fswatch/releases/download/1.15.0/fswatch-1.15.0.tar.gz
RUN tar -xvzf fswatch-1.15.0.tar.gz
WORKDIR /root/fswatch-1.15.0
RUN ./configure
RUN make
RUN make install
WORKDIR /anonym

FROM golang:alpine as builder
COPY . /anonym
WORKDIR /anonym
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .

#Production environment
FROM scratch as prod
COPY --from=builder /anonym/config/config.yml /anonym/config/config.yml
COPY --from=builder /anonym/anonym /anonym/anonym
WORKDIR /anonym

