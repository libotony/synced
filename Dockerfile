# Build thor in a stock Go builder container
FROM golang:alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git
WORKDIR  /go/synced
COPY . /go/synced
RUN make

# Pull built binary into a second stage container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/synced/bin/synced /usr/local/bin/

EXPOSE 8000
ENTRYPOINT ["synced"]