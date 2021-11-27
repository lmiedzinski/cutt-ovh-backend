FROM golang:1.17 AS builder
ADD . /app/src
WORKDIR /app/src
RUN go test ./cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app/main cmd/main.go

FROM alpine:latest
COPY --from=builder /app/main ./main
RUN chmod +x ./main
ENTRYPOINT ["./main"]