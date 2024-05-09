# Этап 1: Сборка приложения
FROM node:20-alpine as build

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Устанавливаем зависимости
COPY package*.json ./

# Копируем исходный код приложения в контейнер
COPY . .
RUN npm install
RUN npm i -g prisma
RUN prisma generate

# Собираем статические файлы приложения, если это необходимо
RUN npm run build

# Этап 2: Production образ
FROM node:20-alpine as production

# Устанавливаем переменные окружения, если они вам необходимы
ENV NODE_ENV=production
ENV FASTIFY_ADDRESS=0.0.0.0
ENV FASTIFY_PORT=5001

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем зависимости из предыдущего этапа
COPY --from=build /app/package*.json ./
COPY --from=build /app/node_modules ./node_modules

# Копируем исходный код приложения
COPY --from=build /app/dist ./

# Открываем порт, который используется вашим приложением
EXPOSE 5001

# Запускаем Fastify приложение
CMD ["node", "index.js"]
