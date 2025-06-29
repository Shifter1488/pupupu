# Используем официальный образ Go для сборки
FROM golang:1.20-alpine AS builder

# Устанавливаем необходимые пакеты (sqlite и т.п.)
RUN apk add --no-cache git

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем исходники
COPY . .

# Собираем исполняемый файл
RUN go build -o parcel-app

# Финальный образ — минимальный alpine с CA сертификатами (по необходимости)
FROM alpine:latest

# Установка sqlite (если используется sqlite3 напрямую)
RUN apk add --no-cache sqlite

# Копируем бинарник из стадии сборки
COPY --from=builder /app/parcel-app /parcel-app

# копируем возможные файлы БД (если есть начальные)
COPY tracker.db .

# Открываем порт, если ваше приложение слушает (в данном примере порт не показан)
# EXPOSE 8080

# Команда запуска
CMD ["/parcel-app"]
