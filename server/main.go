package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//Users структура для парсинга json
type Users struct {
	Id        string `json:id`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//go get -u github.com/gorilla/mux
var (
	db *sql.DB
)

//PORT Порт
const PORT string = ":8080"
const DB_CONNECT_STRING =
	"host= 172.17.0.2 port=5432 user=postgres  password= docker dbname=clientserver sslmode=disable"

func main() {
	DBconnect()
	//Run server and routes
	r := mux.NewRouter()

	//Получить всех пользователей
	r.HandleFunc("/user", handler).Methods("GET")
	//Создать пользователя
	r.HandleFunc("/user", createUsers).Methods("POST")
	//Удалить пользователя
	r.HandleFunc("/user/{id:[0-9]+}", deleteUser).Methods("DELETE")
	//Получить пользователя по id
	r.HandleFunc("/user/{id:[0-9]+}", handlerUser).Methods("GET")
	//Обновить пользователя
	r.HandleFunc("/user/{id:[0-9]+}", updateUser).Methods("PUT")

	log.Println("Server up and run on port " + PORT)
	log.Fatal(http.ListenAndServe(PORT, r))

}

// func handler | Use for get all users
func testall(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world!"))
}
func handler(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT * FROM users")
	PanicOnErr(err)
	defer rows.Close()

	users := make([]*Users, 0)

	for rows.Next() {
		us := new(Users)
		err = rows.Scan(&us.Id, &us.Name, &us.FirstName, &us.LastName)
		PanicOnErr(err)
		users = append(users, us)
	}
	PanicOnErr(err)
	productsJson, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productsJson)
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := Users{}

	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO users (username, first_name, last_name) VALUES ($1, $2, $3)", user.Name, user.FirstName, user.LastName)
	PanicOnErr(err)

	fmt.Println(result)

}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/user/"):]
	index, _ := strconv.ParseInt(id, 10, 0)

	row := db.QueryRow("SELECT * FROM users WHERE id = $1", index)

	us := new(Users)

	err := row.Scan(&us.Id, &us.Name, &us.FirstName, &us.LastName)
	PanicOnErr(err)

	productsJson, _ := json.Marshal(us)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productsJson)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	//deleteUsers
	id := r.URL.Path[len("/user/"):]
	index, _ := strconv.ParseInt(id, 10, 0)

	result, err := db.Exec("DELETE FROM users WHERE id = $1", index)
	PanicOnErr(err)

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User %s delete successfully (%d row affected)\n", id, rowsAffected)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	//update
	id := r.URL.Path[len("/user/"):]
	index, _ := strconv.ParseInt(id, 10, 0)

	decoder := json.NewDecoder(r.Body)
	user := Users{}
	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := db.Exec("UPDATE users SET username = $1, first_name = $2, last_name = $3  WHERE id = $4", user.Name, user.FirstName, user.LastName, index)
	PanicOnErr(err)

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User %s update successfully (%d row affected)\n", id, rowsAffected)

}

//DBconnect run and connect DB
func DBconnect() {
	var err error
	db, err = sql.Open("postgres", DB_CONNECT_STRING)
	PanicOnErr(err)

	err = db.Ping()
	PanicOnErr(err)
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
