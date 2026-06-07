# Stage 1: Build Stage
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final Run Stage
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/swagger.html .
COPY --from=builder /app/api-contract.yaml .
COPY --from=builder /app/db ./db
EXPOSE 8888
CMD ["./main"]