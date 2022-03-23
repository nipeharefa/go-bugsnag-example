FROM golang:1.18-alpine AS builder
RUN apk update && apk add build-base
RUN apk add --no-cache \
        ca-certificates tzdata openssl

ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 CGO_LDFLAGS="-static" go build -ldflags='-w -s ' -o build/auth *.go

FROM alpine
RUN apk update && apk add --no-cache ca-certificates
RUN rm -vrf /var/cache/apk/*
COPY --from=builder /app/build /app
CMD ["./app/auth"]