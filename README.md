# CLIENT-SERVER APP
## Simple realisation of application rest api on Golang..

---
## Task
Требуется реализовать клиент-серверное приложение на языке программирования Golang. Протокол общения rest api.

Сервер необходимо построить на базе пакета https://github.com/gorilla/mux и он должен содержать логику следующего роутинга:
1. Добавить пользователя
2. Удалить пользовалятеля
3. Получить пользователя
4. Получить всех пользователей
5. Обновить пользователя

Для проверки запросов можно использовать postman или curl.
Данные хранить в postgresql. Сервер и postgresql запускать в docker.

Консольный клиент(CLI) реализовать с использованием пакета https://github.com/jawher/mow.cli.

Команды CLI должны уметь работать с реализованными роутами сервера.
Пример:

cli create user -f <path_to_json>

path_to_json - путь к json файлу где описана структура пользователя необходимая для его создания

{
    “username”: “demo”,
    “first_name”:“Ivan”,
    “last_name”:“Ivanov”
}

Написать тесты и опубликовать исходный код на github.

---
## RUN App
### Команды для запуска контенера сервера
docker build -t client-server .

docker run --publish 6060:8080 --name app --rm client-server

###Команды для запуска сервера БД

docker build -t db_serv .

docker run -d db_serv
