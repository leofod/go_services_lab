# mai_lab

## Запуск
```
docker-compose build && docker-compose up
```

## Миграция БД
```
docker run -v $(pwd)/migrations:/migrations --network go_services_lab_go_app migrate/migrate -path=/migrations/ -database postgres://postgres:qweasd@postgres_container:5432/postgres?sslmode=disable up 2
```

## PgAdmin
По пути `http://localhost:5050`

## Сервис пользователей

- Сервис Users
- Сервис позволяет создавать/получать аккаунт пользователя
- Сервис должен хранить все данные в кеше

## Сервис заказов

- Сервис Orders
- Сервис позволяет создавать/хранить/получать заказы
- Сервис должен хранить все данные в кеше

## Ограничения

- ttp server
- Server предоставляет API методы для взаимодействия с вашим сервисом
- Server Users слушает 80 порт
- Server Orders слушает 81 порт
- Для всех операций использовать method POST
