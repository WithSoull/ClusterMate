# ClusterMate

**ClusterMate** — это RESTful API-сервис для управления пользовательскими кластерами и ролями, позволяющий автоматизировать управление доступом и упрощать мониторинг состояния кластеров.

## Функциональность

- Управление пользователями и их ролями в кластерах
- Создание, изменение и удаление кластеров
- Настройка ролей и прав доступа для пользователей
- Поддержка отказоустойчивости кластеров и функций мониторинга
- Версионированный API для упрощения расширяемости

## Стек технологий

- **Go** — основной язык разработки
- **Chi** — маршрутизация и middleware для создания RESTful API
- **Environment Configuration** — поддержка конфигурации через `.env` файл

## Структура проекта

```plaintext
ClusterMate/
├── cmd/
│   └── server/               # Точка входа приложения
│       └── main.go
├── internal/
│   ├── api/                  # Пакет API с версионированием
│   │   └── v1/
│   │       ├── handlers.go   # Обработчики API
│   │       └── routers.go    # Маршруты API
│   ├── config/               # Конфигурация приложения
│   ├── db/                   # Подключение к базе данных
│   └── models/               # Определения моделей данных
├── go.mod
├── go.sum
└── README.md
```


## API Эндпоинты

### Примеры основных маршрутов

- **GET /api/v1/users** — получение списка пользователей
- **POST /api/v1/users** — создание нового пользователя
- **GET /api/v1/clusters** — получение списка кластеров
- **POST /api/v1/clusters** — создание нового кластера
- **GET /healthcheck** — проверка доступности сервиса

Подробную информацию о каждом эндпоинте можно найти в [документации API](/docs).
