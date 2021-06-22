FROM golang:1.16.2 AS builder

RUN apt-get update && apt-get install upx -y
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
COPY cert.pem cert.pem
COPY key.pem key.pem

# Build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"
RUN go mod download && go build -o myweb main.go && upx myweb

FROM alpine:3.9.2
COPY --from=builder /workspace/myweb .
COPY --from=builder /workspace/cert.pem .
COPY --from=builder /workspace/key.pem .
ENTRYPOINT ["/myweb"]
