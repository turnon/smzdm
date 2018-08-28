FROM golang:alpine as builder

WORKDIR $GOPATH/src/github.com/turnon/smzdm
COPY *.go ./

RUN apk add --no-cache git \
    && git clone https://github.com/golang/net $GOPATH/src/golang.org/x/net \
    && go get ./... \
    && go build -o /smzdm \
    && apk del git

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /smzdm .
RUN chmod +x /root/smzdm
RUN apk update && apk add tzdata

ENTRYPOINT ["/root/smzdm"]
