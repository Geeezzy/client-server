package main

import (
	"fmt"
	"log"
	"net/http"
)

//PORT Порт
const PORT string = "8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request ", r.URL.Path)
		defer r.Body.Close()

		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "Hello world!")
		case http.MethodPost:
		case http.MethodPut:
		case http.MethodDelete:
		}
	})
	log.Println("Server up and run on port ", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
