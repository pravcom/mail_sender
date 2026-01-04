.PHONY: docker docker-build docker-up docker-down docker-logs docker-restart docker-shell docker-ls docker-tree

# Сборка и запуск контейнера
docker:
	docker compose up -d --build

# Только сборка образа
docker-build:
	docker compose build

# Запуск контейнера
docker-up:
	docker compose up -d

# Остановка контейнера
docker-down:
	docker compose down

# Просмотр логов
docker-logs:
	docker compose logs -f -t app

# Перезапуск контейнера
docker-restart:
	docker compose restart app

# Просмотр статуса
docker-ps:
	docker compose ps

# Интерактивный shell в контейнере
docker-shell:
	docker compose exec app sh

# Просмотр файлов в рабочей директории
docker-ls:
	docker compose exec app ls -la

# Просмотр структуры директорий (если установлен tree)
docker-tree:
	docker compose exec app sh -c "find . -type d | head -20"