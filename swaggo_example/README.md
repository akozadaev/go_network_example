# Swagger / OpenAPI - пример REST API с документацией

Пример REST API для управления пользователями с автогенерируемой документацией Swagger (OpenAPI) на основе аннотаций в коде Go.

## Описание

Приложение представляет собой REST API с CRUD-операциями для сущности `User`. Спецификация API описывается swag-аннотациями в коде; инструмент `swag` генерирует файлы `docs/docs.go`, `docs/swagger.json` и `docs/swagger.yaml`, которые используются библиотекой `http-swagger` для отображения интерактивной документации Swagger UI.

## API эндпоинты

| Метод | Путь | Описание |
|-------|------|----------|
| GET | /users | Список всех пользователей |
| GET | /users/{id} | Пользователь по ID |
| POST | /users | Создание пользователя (тело: `{"name": "..."}`) |

Swagger UI доступен по адресу: `http://localhost:8080/swagger/index.html`

## Установка зависимостей

```bash
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/http-swagger
```

Установите `swag` глобально (для генерации документации):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Генерация документации

После изменения аннотаций в коде выполните:

```bash
swag init -g main.go -o ./docs
```

- `-g main.go` - файл с описанием сервера и роутов
- `-o ./docs` - каталог для сгенерированных файлов

## Запуск

```bash
go run main.go
```

- REST API: http://localhost:8080
- Swagger UI: http://localhost:8080/swagger/index.html

## Примеры запросов

```bash
# Получить всех пользователей
curl http://localhost:8080/users

# Получить пользователя по ID
curl http://localhost:8080/users/1

# Создать пользователя
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Charlie"}' http://localhost:8080/users
```
