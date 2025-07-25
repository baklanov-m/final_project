### О проекте

Веб-сервер — планировщик задач (TODO-лист) с поддержкой повторяющихся задач. Задачи содержат дату, заголовок и комментарий. При выполнении повторяющихся задач дата обновляется согласно правилу, обычные задачи удаляются.

### Функции API

- Добавление, получение, удаление задач
- Получение и изменение параметров задачи
- Пометка задачи как выполненной

### Структура проекта

```
pkg/
 ├─ api/       # API-обработчики
 ├─ db/        # Работа с БД
 ├─ nextdate/  # Логика повторений
 └─ server/    # Запуск сервера
tests/         # Тесты
web/           # Фронтенд (css, js)
```

### Технологии

Go (net/http), SQLite, HTML/CSS/JS

### База данных

Таблица `scheduler` с полями: `id`, `date` (YYYYMMDD), `title`, `comment`, `repeat` (правила повторений).

### Запуск

```bash
go mod tidy
go run main.go
```

Открыть http://localhost:7540

Использование переменных окружения:

```bash
TODO_PORT=8080 TODO_DBFILE="my_tasks.db" go run main.go
```

### Тесты

```bash
go test ./tests
go test -run ^TestApp$ ./tests
```

### Документация

```bash
go install golang.org/x/tools/cmd/godoc@latest
go doc -http=:6060
```

Открыть http://localhost:6060/pkg/todo/

### Особенности

- Порт и файл БД настраиваются через `TODO_PORT` и `TODO_DBFILE`.
- В API `/api/tasks` реализован поиск по тексту и дате (формат DD.MM.YYYY).

---
