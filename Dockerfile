FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/internal

FROM alpine:latest

# Устанавливаем базу данных временных зон для работы с time.LoadLocation
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]