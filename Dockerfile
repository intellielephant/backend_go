FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
   GOPROXY=https://goproxy.cn,direct \
   GIN_MODE=release

# ENV mysqlAddress=172.17.0.1

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY ./ ./

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]