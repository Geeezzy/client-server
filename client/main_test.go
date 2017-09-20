package main

import (
	"testing"
	"net/http/httptest"
	"net/http"


	//"log"
	//"io/ioutil"
	"fmt"
	"log"
	"io/ioutil"
)

func TestGetAction(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	ts.URL = DEFAULT_HOST

}

/*func TestCreateAction(t *testing.T) {

}
/*func TestDeleteAction(t *testing.T) {

}
func TestUpdateAction(t *testing.T) {

}*/
