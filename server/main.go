package main

import (
	"database/sql"
	"encoding/json"
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

func hundler(w http.ResponseWriter, r *http.Request) {

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

func main() {
	DBconnect()
	//Run server and routes
	r := mux.NewRouter()
	r.HandleFunc("/", hundler).Methods("GET")
	log.Println("Server up and run on port " + PORT)
	log.Fatal(http.ListenAndServe(PORT, r))

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
