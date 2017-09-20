package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)
func TestGetUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	resp := GetUsers(server.URL)

	if resp != "200 OK" {
		t.Error("BadRequest")
	}
}
func TestGetUserById(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	resp := GetUserById("2", server.URL)

	if resp != "200 OK" {
		t.Error("BadRequest")
	}
}
func TestDeleteUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	resp := DeleteUser("2", server.URL)

	if resp != "200 OK" {
		t.Error("BadRequest")
	}
}
func TestCreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	resp := CreateUser("files/addUser.json", server.URL)

	if resp != "200 OK" {
		t.Error("BadRequest")
	}
}
func TestUpdateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	resp := UpdateUser("files/updateUser.json", "2", server.URL)

	if resp != "200 OK" {
		t.Error("BadRequest")
	}
}
