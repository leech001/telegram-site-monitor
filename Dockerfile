FROM golang:latest AS builder
WORKDIR /build
COPY app .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o monitor


FROM alpine:latest
RUN apk add --no-cache bash
WORKDIR /app
COPY --from=builder /build/config.yaml .
COPY --from=builder /build/monitor .
CMD sh -c "./monitor"