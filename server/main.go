package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//go get -u github.com/gorilla/mux
var (
	db *sql.DB
)

//PORT Порт
const PORT string = ":8080"

func hundler(w http.ResponseWriter, r *http.Request) {
	var username string
	rows, err := db.Query("SELECT username FROM users")
	PanicOnErr(err)
	for rows.Next() {

		err = rows.Scan(&username)
		PanicOnErr(err)
		//fmt.Println("rows.Next username: ", username)
	}
	rows.Close()

	w.Write([]byte("rows.Next username: "))
	w.Write([]byte(username))
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
