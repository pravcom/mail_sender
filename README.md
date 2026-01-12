### Как поменять настройки
- Подключиться к серверу по SSH через putty
- Перейти в cd mail_sender
- sudo nano .env
- docker ps -a посмотреть id контейнера mail_sender если существует
- docker rm "id" удалить старый образ
- docker-compose up сгенерировать новый образ
