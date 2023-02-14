FROM golang:1.18 as builder

WORKDIR /code
ADD go.mod .
ADD go.sum .
ADD . .
RUN go mod download



RUN go build -o server ./cmd/server

FROM ubuntu:20.04
WORKDIR /code
COPY --from=builder /code/server .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD "/code/server"
