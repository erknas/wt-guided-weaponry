FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/wt-guided-weaponry/main.go 

EXPOSE 8000

CMD ["./main"]