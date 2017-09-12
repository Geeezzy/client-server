package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func main() {
	DBconnect()
	//Run server and routes
	r := mux.NewRouter()

	//Получить всех пользователей
	r.HandleFunc("/getallusers", handler).Methods("GET")
	//Создать пользователя
	r.HandleFunc("/createuser", createUsers).Methods("POST")
	//Удалить пользователя
	r.HandleFunc("/deleteuser", deleteUser).Methods("DELETE")
	//Получить пользователя по id
	r.HandleFunc("/getuser", handlerUser).Methods("GET")
	//Обновить пользователя
	r.HandleFunc("/update", updateUser).Methods("PUT")

	log.Println("Server up and run on port " + PORT)
	log.Fatal(http.ListenAndServe(PORT, r))

}

// func handler | Use for get all users
func handler(w http.ResponseWriter, r *http.Request) {

	//users := []Users{}
	//var username string
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
	name := r.FormValue("name")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")

	//Check for empty values
	if name == "" || first_name == "" || last_name == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	//INSERT
	result, err := db.Exec("INSERT INTO users VALUES($1, $2, $3)", name, first_name, last_name)
	PanicOnErr(err)

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User %s created successfully (%d row affected)\n", name, rowsAffected)

}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)

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

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	//update
}

//DBconnect run and connect DB
func DBconnect() {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=clienserver sslmode=disable")
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
