FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o echo-server main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/echo-server .

EXPOSE 8080

CMD ["./echo-server"]
