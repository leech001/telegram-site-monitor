FROM golang:latest AS builder
WORKDIR /build
COPY app .
RUN go get main
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-s -w -extldflags "-static"' -o monitor


FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/monitor .
CMD sh -c "./monitor"