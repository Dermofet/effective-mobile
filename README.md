# Effective Mobile Test API 

## Требования

Проект разрабоатывался на **Golang** версии 1.21.5.

## Установка

1. Клонируйте репозиторий с проектом:
   ```bash
   git clone https://github.com/Dermofet/effective-mobile
   ```
   
2. Перейдите в директорию проекта:
    ```bash
    cd effective-mobile
    ```
    
3. Установите зависимости:
    ```bash
    go mod download
    ```
    
## Запуск

Для запуска Docker контейнеров выполните следующие команды:

```bash
docker compose -f ./dev/docker-compose.yml up -d --build
```

Для обновления документации Swagger выполните следующую команду:
```bash
swag init -g .\cmd\effective-mobile-test\main.go --parseInternal
```

Для просмотра документации Swagger перейдите по ссылке:

**http://localhost:8000/swagger/index.html**

## Конфигурация

Для конфигурации проекта используется файл **.env**.

#### Переменные окружения

Для настройки приложения:
- APP_NAME
- APP_VERSION
- LOG_LEVEL
- DEBUG
- PATH_LOG

Для настройки работы с внешним API:
- SERVICE_URL - необходимо вставить адрес внешнего API, без него проект не будет работать

Для настройки http сервера:
- HTTP_HOST
- HTTP_PORT

Для настройки базы данных:
- DB_HOST
- DB_PORT
- DB_NAME
- DB_USER
- DB_PASS
- DB_SSLMODE
