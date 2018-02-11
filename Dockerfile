FROM golang:1.9-alpine AS backend
RUN apk add --no-cache git
WORKDIR "/go/src/ip-notify"
COPY . .
RUN dep ensure && go build -o /tmp/ip-notify

FROM alpine:3.5
LABEL maintainer=benkris1@126.com
COPY --from=backend /tmp/ip-notify /usr/local/bin
ENTRYPOINT [ "ip-notify" ]