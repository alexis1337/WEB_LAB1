# Dockerfile для frontend
# Используем Node.js для сборки
FROM node:18 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY package.json package-lock.json ./
RUN npm install

COPY . .

# Собираем фронтенд
RUN npm run build

FROM nginx:stable-alpine

COPY --from=builder /app/dist /usr/share/nginx/html

# Открываем порт 80
EXPOSE 80