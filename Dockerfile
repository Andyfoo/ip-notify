FROM golang:1.9-alpine
LABEL maintainer=benkris1@126.com
RUN apk add --no-cache git && go get -u github.com/golang/dep/cmd/dep
RUN mkdir -p /go/src/ip-notify

WORKDIR "/go/src/ip-notify"
COPY . .
RUN dep ensure && go build -o /tmp/ip-notify
RUN cp /tmp/ip-notify /usr/local/bin
ENTRYPOINT [ "ip-notify" ]

