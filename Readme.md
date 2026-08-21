## 📋 Оглавление

1. Контекст проекта
    
2. Структура проекта
    
3. Разработка (локальный цикл)
    
4. Docker (сборка и запуск)
    
5. Полная шпаргалка команд
    
6. Частые ошибки и решения
    

---

## 1. КОНТЕКСТ ПРОЕКТА

**Название:** TS4 Randomizer Bot  
**Описание:** Telegram-бот для генерации случайных событий в игре Sims 4  
**Технологии:**

- Go 1.26.5
    
- SQLite ([modernc.org/sqlite](https://modernc.org/sqlite))
    
- Docker
    
- Telegram Bot API
    

**Режимы работы:**

- **Локально** — для разработки и тестирования (на Windows)
    
- **Docker** — для изолированного запуска (эмуляция Linux-среды)
    

---

## 2. СТРУКТУРА ПРОЕКТА

text

D:\TS4RandomizerBot\
├── data/
│   └── events.db              # База данных SQLite
├── db/
│   ├── database.go            # Подключение и инициализация БД
│   ├── model.go               # Структуры данных
│   └── repository.go          # CRUD-операции
├── migrations/
│   └── schema.sql             # Схема БД
├── .dockerignore              # Исключения для Docker
├── .env                       # Переменные окружения
├── .gitignore                 # Исключения для Git
├── Dockerfile                 # Инструкция для сборки Docker-образа
├── go.mod                     # Go-модуль
├── go.sum                     # Контрольные суммы зависимостей
├── main.go                    # Точка входа
└── Makefile                   # Автоматизация команд

---

## 3. ЦИКЛ РАЗРАБОТКИ (ЛОКАЛЬНО)

### 📝 Шаг 1: Проверка изменений

**Где:** Git Bash (в папке проекта `D:\TS4RandomizerBot`)

bash

git status

**Что показывает:** Какие файлы изменены, добавлены или удалены.

---

### 📝 Шаг 2: Добавление изменений в Git

bash

git add .

**Что делает:** Добавляет все изменённые файлы в индекс для коммита.

---

### 📝 Шаг 3: Создание коммита

bash

git commit -m "Описание изменений"

**Что делает:** Сохраняет состояние проекта с комментарием.

---

### 📝 Шаг 4: Отправка на GitHub

bash

git push

**Что делает:** Отправляет изменения в удалённый репозиторий.

---

### 📝 Шаг 5: Проверка работы (локально)

bash

go run .

**Что делает:** Запускает бота локально для тестирования.

**Остановка:** `Ctrl+C`

---

## 4. DOCKER (СБОРКА И ЗАПУСК)

### 🐳 Шаг 1: Сборка Docker-образа

**Где:** Git Bash

bash

make build

**Что делает:** Собирает образ `ts4-bot` на основе `Dockerfile`.

---

### 🐳 Шаг 2: Запуск контейнера

bash

make run

**Что делает:** Запускает контейнер с ботом, монтирует папку `data`.

**Остановка:** `Ctrl+C`

---

### 🐳 Шаг 3: Сборка + запуск одной командой

bash

make dev

**Что делает:** Выполняет `make build` и сразу `make run`.

---

### 🐳 Шаг 4: Остановка контейнера (в фоне)

bash

make stop

**Что делает:** Останавливает контейнер `my-bot`.

---

### 🐳 Шаг 5: Просмотр логов

bash

make logs

**Что делает:** Показывает логи запущенного контейнера.

---

### 🐳 Шаг 6: Вход в контейнер (отладка)

bash

make shell

**Что делает:** Открывает командную оболочку внутри контейнера.

---

### 🐳 Шаг 7: Очистка (удаление образа и контейнера)

bash

make clean

**Что делает:** Удаляет контейнер, образ и освобождает место.

---

## 5. ПОЛНАЯ ШПАРГАЛКА КОМАНД

### 🔧 Git команды

|Команда|Где|Что делает|
|---|---|---|
|`git status`|Git Bash|Показывает состояние репозитория|
|`git add .`|Git Bash|Добавляет все изменения|
|`git commit -m "..."`|Git Bash|Создаёт коммит|
|`git push`|Git Bash|Отправляет на GitHub|
|`git pull`|Git Bash|Забирает изменения с GitHub|

---

### 🐳 Docker команды

|Команда|Где|Что делает|
|---|---|---|
|`docker build -t ts4-bot .`|Git Bash|Собирает образ|
|`docker run -it --rm --name my-bot -v D:/TS4RandomizerBot/data:/root/data ts4-bot`|Git Bash|Запускает контейнер|
|`docker stop my-bot`|Git Bash|Останавливает контейнер|
|`docker logs -f my-bot`|Git Bash|Показывает логи|
|`docker exec -it my-bot sh`|Git Bash|Заходит внутрь контейнера|
|`docker ps`|Git Bash|Список запущенных контейнеров|
|`docker ps -a`|Git Bash|Все контейнеры (включая остановленные)|
|`docker rm my-bot`|Git Bash|Удаляет контейнер|
|`docker rmi ts4-bot`|Git Bash|Удаляет образ|

---

### 🛠️ Make команды (упрощённые)

|Команда|Где|Что делает|
|---|---|---|
|`make build`|Git Bash|Сборка образа|
|`make run`|Git Bash|Запуск контейнера|
|`make dev`|Git Bash|Сборка + запуск|
|`make stop`|Git Bash|Остановка контейнера|
|`make logs`|Git Bash|Просмотр логов|
|`make shell`|Git Bash|Вход в контейнер|
|`make clean`|Git Bash|Очистка|

---

### 🗄️ SQLite команды

|Команда|Где|Что делает|
|---|---|---|
|`sqlite3 data/events.db ".read migrations/schema.sql"`|PowerShell|Создаёт БД из схемы|
|`sqlite3 data/events.db`|PowerShell|Открывает БД в консоли|
|`.tables`|SQLite|Показывает таблицы|
|`.schema events`|SQLite|Показывает структуру таблицы|
|`.exit`|SQLite|Выход из SQLite|

---

## 6. ЧАСТЫЕ ОШИБКИ И РЕШЕНИЯ

### ❌ Ошибка: `make: command not found`

**Причина:** GNU Make не установлен.

**Решение:**

bash

choco install make -y

Или установи вручную с [gnuwin32.sourceforge.net](https://gnuwin32.sourceforge.net/packages/make.htm).

---

### ❌ Ошибка: `Binary was compiled with 'CGO_ENABLED=0'`

**Причина:** Используется драйвер `mattn/go-sqlite3`, требующий CGO.

**Решение:**

bash

go get modernc.org/sqlite
go mod tidy

---

### ❌ Ошибка: `open migrations/schema.sql: no such file or directory`

**Причина:** Папка `migrations/` не скопирована в контейнер.

**Решение:** Добавь в Dockerfile:

dockerfile

COPY --from=builder /app/migrations ./migrations

---

### ❌ Ошибка: `go.mod requires go >= 1.26.5 (running go 1.22.12)`

**Причина:** Несовпадение версий Go.

**Решение:** Исправь `go.mod`:

text

go 1.22

**ИЛИ** обнови образ в Dockerfile:

dockerfile

FROM golang:1.26.5-alpine AS builder

---

### ❌ Ошибка: `-v /data:/root/data` — БД не подтягивается

**Причина:** Относительный путь не работает в Windows.

**Решение:** Используй абсолютный путь:

makefile

-v D:/TS4RandomizerBot/data:/root/data

---

### ❌ Ошибка: `Virtualization support not detected`

**Причина:** Виртуализация отключена в BIOS.

**Решение:**

1. Войти в BIOS (F2, Delete, F10)
    
2. Включить Intel VT-x / AMD-V
    
3. Включить Virtual Machine Platform в компонентах Windows
    

---

## 7. ПОЛЕЗНЫЕ ССЫЛКИ

|Ресурс|Ссылка|
|---|---|
|Docker Desktop|[docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop/)|
|SQLite Studio|[sqlitestudio.pl](https://sqlitestudio.pl/)|
|Go|[go.dev](https://go.dev/)|
|Chocolatey|[chocolatey.org](https://chocolatey.org/)|
|Telegram Bot API|[core.telegram.org/bots/api](https://core.telegram.org/bots/api)|

---

## 8. ЕЖЕДНЕВНЫЙ ЦИКЛ РАБОТЫ (КРАТКО)

|Шаг|Команда|Где|
|---|---|---|
|1|`git add . && git commit -m "fix" && git push`|Git Bash|
|2|`make dev`|Git Bash|
|3|Тестируешь бота в Telegram|—|
|4|`Ctrl+C` (остановка)|Git Bash|