# Build
ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod /api
COPY go.sum /api
COPY . /api

RUN go mod download
RUN go mod tidy
RUN cd /api && go build -o /api/executable .

# Run
FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api/executable .

EXPOSE ${PORT}

ENTRYPOINT ["./executable"]