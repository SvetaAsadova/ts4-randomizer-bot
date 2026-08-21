# Makefile для автоматизации сборки и запуска бота

# Переменные
IMAGE_NAME = ts4-bot
CONTAINER_NAME = my-bot
DATA_PATH = $(shell pwd)/data

# Цель по умолчанию
.PHONY: help
help:
	@echo "Доступные команды:"
	@echo "  make build    - Собрать Docker-образ"
	@echo "  make run      - Запустить контейнер с ботом"
	@echo "  make stop     - Остановить контейнер"
	@echo "  make restart  - Перезапустить контейнер"
	@echo "  make shell    - Зайти внутрь контейнера (отладка)"
	@echo "  make logs     - Посмотреть логи бота"
	@echo "  make clean    - Удалить образ и контейнер"

# Сборка Docker-образа
.PHONY: build
build:
	@echo "📦 Сборка Docker-образа..."
	docker build -t $(IMAGE_NAME) .

# Запуск контейнера
.PHONY: run
run:
	@echo "🚀 Запуск бота..."
	docker run -it --rm --name my-bot -v D:/TS4RandomizerBot/data:/root/data ts4-bot

# Запуск в фоне (без логов)
.PHONY: start
start:
	@echo "🚀 Запуск бота в фоне..."
	docker run -d --rm --name $(CONTAINER_NAME) -v $(DATA_PATH):/root/data $(IMAGE_NAME)
	@echo "✅ Бот запущен в фоне. Используй 'make logs' для просмотра логов."

# Остановка контейнера
.PHONY: stop
stop:
	@echo "🛑 Остановка контейнера..."
	-docker stop $(CONTAINER_NAME)

# Перезапуск
.PHONY: restart
restart: stop build run

# Зайти внутрь контейнера
.PHONY: shell
shell:
	@echo "🔧 Заходим внутрь контейнера..."
	docker exec -it $(CONTAINER_NAME) sh

# Просмотр логов
.PHONY: logs
logs:
	@echo "📋 Логи бота:"
	docker logs -f $(CONTAINER_NAME)

# Очистка
.PHONY: clean
clean:
	@echo "🧹 Очистка..."
	-docker stop $(CONTAINER_NAME)
	-docker rm $(CONTAINER_NAME)
	-docker rmi $(IMAGE_NAME)
	@echo "✅ Очистка завершена."

# Быстрая сборка + запуск (одной командой)
.PHONY: dev
dev: build run