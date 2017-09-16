##Команды для запуска контенера сервера


docker build -t client-server .
docker run --publish 6060:8080 --name app --rm client-server

---

##Команды для запуска сервера БД

docker build -t db_serv .
docker run -d db_serv
