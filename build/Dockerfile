ARG TARGET

FROM golang:1.15-alpine3.12 AS builder

RUN apk update &&  apk add --no-cache \
    git \
    ca-certificates \
    gcc    \
    libc-dev    \
    pkgconfig

ARG TARGET
ENV GO111MODULE=on

COPY . /mygo/
WORKDIR /mygo/cmd/$TARGET
RUN GOOS=linux GOARCH=amd64 go build -o $TARGET


FROM alpine:latest
RUN apk update && apk add --no-cache \
    curl    \
    ca-certificates  \
    && rm -rf /var/cache/apk/*  \
    && mkdir /app

ARG TARGET
COPY --from=builder /mygo/cmd/$TARGET /$TARGET

RUN addgroup -g 1000 nobody \
    &&  adduser -D -u 1000 -G nobody nobody  \
    && chown -R nobody:nobody /$TARGET

USER       nobody
ENTRYPOINT ["${TARGET}"]
#CMD ["run"]