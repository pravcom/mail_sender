### Как поменять настройки
- Подключиться к серверу по SSH через putty
- Перейти в cd mail_sender
- sudo nano .env
- docker ps -a посмотреть id контейнера mail_sender если существует
- docker rm "id" удалить старый образ
- docker-compose up сгенерировать новый образ
## Если docker-compose up не работает
- sudo apt install docker-compose
## Создание нового пользователя
- sudo adduser newuser
- sudo usermod -aG sudo newuser
Переключиться на нового пользователя
- su - newuser
- sudo whoami
## Убрать обязательный ввод пароля
Откройте файл конфигурации
- sudo visudo
Добавьте новое правило
- username ALL=(ALL) NOPASSWD: ALL
