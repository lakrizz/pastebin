FROM golang:alpine AS builder

EXPOSE 8000/tcp

ENTRYPOINT ["pastebin"]

RUN \
    apk add --update git && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /go/src/pastebin
WORKDIR /go/src/pastebin

COPY . /go/src/pastebin

RUN go get -v -d
RUN go get github.com/GeertJohan/go.rice/rice
RUN GOROOT=/go rice embed-go
RUN go install -v


FROM alpine

EXPOSE 8000/tcp
ENTRYPOINT ["pastebin"]

COPY --from=builder /go/bin/pastebin /bin/pastebin


