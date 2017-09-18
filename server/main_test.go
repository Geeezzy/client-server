package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"encoding/json"
)
type User struct {
	Name string
	FirstName string
	LastName string
}

func TestGetUsers(t *testing.T){

	req, err := http.NewRequest("GET", "/user", nil)

	if err != nil {

		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetUsers)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {

		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetUserById(t *testing.T) {
	req,err := http.NewRequest("GET", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	handler := http.HandlerFunc(GetUserById)
	handler.ServeHTTP(res,req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status,http.StatusOK)
	}
}

func TestCreateUsers(t *testing.T) {

	u := User{
		Name:      "test2",
		FirstName: "lol2",
		LastName:  "kek2",
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	req, err := http.NewRequest("POST", "/user", b)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUsers)
	handler.ServeHTTP(res,req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status,http.StatusOK)
	}

}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/user/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status,http.StatusOK)
	}

}

func TestUpdateUser(t *testing.T) {
	u := User{
		Name:      "test2",
		FirstName: "lol2",
		LastName:  "kek2",
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	req, err := http.NewRequest("PUT", "/user/1", b)
	if err != nil{
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateUser)
	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status,http.StatusOK)
	}
}

