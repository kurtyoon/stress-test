FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o build/stress-test cmd/loadtest/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/build/stress-test .
COPY --from=builder /app/web/dashboard.html ./web/

EXPOSE 8080

ENTRYPOINT ["./stress-test"]
CMD ["-url", "", "-rps", "0", "-duration", "0"]