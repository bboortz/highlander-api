#
# stage 1
# 
FROM golang:latest as build

WORKDIR /
RUN mkdir -p /go/src/github.com/bboortz/highlander-api
WORKDIR /go/src/github.com/bboortz/highlander-api
COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

RUN go get github.com/ghodss/yaml && \
  CGO_ENABLED=0 GOOS=linux go build -a -tags "netgo static_build" -installsuffix netgo -ldflags "-w -s" -o highlander cmd/main.go



#
# stage 2
#
FROM scratch
LABEL maintainer "Benjamin Boortz <benjamin.boortz@secure.mailbox.org>"

WORKDIR /
COPY examples /examples
COPY examples/highlander.yaml /
COPY --from=build go/src/github.com/bboortz/highlander-api/highlander /
ENTRYPOINT ["/highlander"]
