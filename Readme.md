# Image Service

[![Linters Status](https://github.com/AITH-Courses/ImageService/actions/workflows/golangci-lint.yml/badge.svg?branch=master)](https://github.com/AITH-Courses/ImageService/actions/workflows/golangci-lint.yml)

## Разработка
### Окружение для разработки
```bash
docker-compose -f dev.docker-compose.yaml up -d
```
### Перезапуск сервиса
```bash
docker-compose -f dev.docker-compose.yaml restart image_service
```
### Пересборка сервиса (если изменился состав библиотек)
```bash
docker-compose -f dev.docker-compose.yaml build
```

### Линтеры
```bash
golangci-lint run
```

### Форматтеры
```bash
gofmt -s -w .
```