FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/health-check main.go

FROM alpine:latest
COPY --from=builder /app/health-check /health-check
CMD ["/health-check"]