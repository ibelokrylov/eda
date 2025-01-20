
# Этап 1: Сборка исполняемого файла
# Используем базовый образ Golang для сборки приложения
FROM golang:1.22.3-alpine as builder

# Установка необходимых пакетов: Git и Hyperscan
RUN apk add --no-cache git cmake g++ libstdc++ libhyperscan-dev

# Создаем рабочую директорию для нашего приложения
WORKDIR /app

# Копирование модульных файлов Go для оптимизации слоев кеша
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование остальных файлов проекта
COPY . .

# Сборка приложения (исполняемого файла)
RUN CGO_ENABLED=0 go build -ldflags='-w -s' -o /main .

# Этап 2: Сборка конечного образа
# Используем небольшой образ Alpine Linux для запуска приложения
FROM alpine:3.16

# Установка CA-сертификатов и Hyperscan runtime library
RUN apk add --no-cache ca-certificates libstdc++ libhyperscan

# Создаем непривилегированного пользователя для запуска приложения (защита)
RUN adduser -D user
USER user

# Копирование исполняемого файла из предыдущего этапа
COPY --from=builder /main /main
COPY .env .env

# Объявление порта, на котором будет работать приложение
EXPOSE 5001

# Запуск приложения
CMD ["/main"]
