# ToDo application
Серверная часть приложения для создания списков задач

##В ходе разработки приложения были получены следующие навыки:

* разработки веб-приложения на Go с дизайном REST API;
* работы с фреймворком gin-gonic/gin;
* внедрения зависимостей;
* конфигурирования приложения с помощью библиотеки spf13/viper и работы с переменными окружения;
* работы с БД Postgresql с использованием библиотеки sqlx;
* написания SQL-запросов

## Запуск приложения
### Docker
```azure
$ docker-compose up todo-app
```
### Terminal
```azure
$ docker run --name=todo_db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
$ docker ps
$ docker exec -it <id>
$ migrate -path ./todo-app/schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
$ go run ./cmd/main.go
```