FROM golang:1.15-alpine AS builder

WORKDIR /src

ENV GOPROXY="https://goproxy.io,direct"
ENV GO111MODULE=on

ADD go.mod go.sum /src/
RUN go mod download

ADD . /src/
RUN go build -o bin/api-server ./cmd/api-server

FROM alpine

WORKDIR /app

COPY --from=builder /src/bin/api-server /app/
RUN ls -al /app

ENTRYPOINT ["/app/api-server"]
