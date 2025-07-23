FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./main.go

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8000
ENTRYPOINT ["./main"]