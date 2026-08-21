# Этап 1: Сборка
FROM golang:1.26.5-alpine AS builder

ENV GOPROXY=https://goproxy.io,direct
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o sims-bot .

# Этап 2: Финальный образ
FROM alpine:latest
WORKDIR /root/

# Копируем бинарник
COPY --from=builder /app/sims-bot .

# Копируем папку с миграциями (ВАЖНО!)
COPY --from=builder /app/migrations ./migrations

# Делаем бинарник исполняемым
RUN chmod +x sims-bot

# Создаём папку для данных
RUN mkdir -p /root/data

CMD ["./sims-bot"]