# Dockerfile для backend
# Используем официальный образ Go для сборки
FROM golang:1.23.2 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Создаём минимальный образ для запуска
FROM alpine:latest

WORKDIR /root/

# Копируем исполняемый файл из стадии сборки
COPY --from=builder /app/main .

# Копируем файлы миграций и другие необходимые ресурсы
COPY migrations ./migrations

# Открываем порт 8080
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
