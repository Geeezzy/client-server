package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestGetUsers(t *testing.T){

	req, err := http.NewRequest("GET", "/user", nil)

	if err != nil {

		t.Fatal(err)
	}

	rr := httptest.NewRecorder()


	handler := http.HandlerFunc(GetUsers)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {

		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
