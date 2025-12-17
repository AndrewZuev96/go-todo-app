# Go To-Do App 🚀

Полноценное Fullstack веб-приложение: список задач с сохранением в базу данных.
Написано в учебных целях для демонстрации навыков Backend-разработки.

## 🛠 Технологический стек

- **Backend:** Golang (net/http, clean architecture)
- **Database:** PostgreSQL
- **Frontend:** HTML/CSS/JS (Vanilla)
- **Infrastructure:** Docker & Docker Compose
- **Config:** Environment variables (.env)

## ✨ Функционал

- [x] Создание задач (POST)
- [x] Просмотр списка (GET)
- [x] Обновление статуса (PUT)
- [x] Удаление задач (DELETE)
- [x] Хранение данных в PostgreSQL

## 🚀 Как запустить (в 1 команду)

Вам понадобится установленный Docker.

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/AndrewZuev96/go-todo-app.git

2. Запуск проекта
    ```bash
    docker-compose up --build

3. Проверка работы

    API Endpoint: http://localhost:8080/tasks
    Frontend: Просто откройте файл index.html в вашем браузере.

👤 Автор

Andrew Zuev

    GitHub: @AndrewZuev96

Created with ❤️ and Golang.
