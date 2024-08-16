FROM golang:1.22.2-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/wt-guided-weaponry/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY .env /app/.env

ARG PORT
EXPOSE $PORT

CMD ["./main"]