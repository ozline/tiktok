FROM golang:1.20 AS builder

ENV TZ Asia/Shanghai
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

RUN mkdir -p /app

WORKDIR /app

ADD . /app
RUN go mod tidy
RUN make build-all


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai
ENV service api

WORKDIR /app
COPY --from=builder /app/output /app/output
COPY --from=builder /app/config /app/config

CMD ["sh", "./docker-entrypoint.sh"]
