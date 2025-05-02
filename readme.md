запуск:

## вариант 1 (без клонирования репозитория через docker hub):
```
docker run -d --name pet-app \
  -p 8080:8080 \
  -p 3000:3000 \
  -e STORAGE_PATH="./storage/storage.json" \
  docker.io/websalad/pet:latest
```

##  вариант 2 (с клонированием репозитория и запуском через docker compose):
- создайте .env в корне по примеру [.env.example](./.env.example)
- выполните:
```
docker-compose up --build
```

### вариант 3 (с клонированием репозитория и запуском через go run):
- запустите приложение
```
STORAGE_PATH="./storage/storage.json" \
go run cmd/main.go
```