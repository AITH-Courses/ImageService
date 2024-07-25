# Image Service

[![Linters Status](https://github.com/AITH-Courses/ImageService/actions/workflows/golangci-lint.yml/badge.svg?branch=master)](https://github.com/AITH-Courses/ImageService/actions/workflows/golangci-lint.yml)

## Разработка
### Окружение для разработки
```bash
docker-compose -f dev.docker-compose.yaml up -d
```

### Запуск приложения
```bash
go run cmd/main.go
```

### Линтеры
```bash
golangci-lint run
```

### Форматтеры
```bash
gofmt -s -w .
```