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

📂 Структура проекта

Проект следует принципам Clean Architecture:
.
├── internal
│   ├── models       # Структуры данных (Data Models)
│   └── storage      # Логика работы с БД (Repository Pattern)
├── index.html       # Frontend клиент
├── main.go          # Точка входа, HTTP Handlers, Config
├── docker-compose.yml
├── Dockerfile
└── .env             # Переменные окружения (не в репозитории)

📡 API Documentation

Метод	URL	Описание	Тело запроса (JSON)
GET     /tasks	Получить все задачи	-
POST    /tasks	Создать задачу	{"title": "...", "completed": false}
PUT	    /tasks	Обновить задачу	{"id": 1, "title": "...", "completed": true}
DELETE	/tasks?id=1	Удалить задачу	

👤 Автор

Andrew Zuev

    GitHub: @AndrewZuev96

Created with ❤️ and Golang.
