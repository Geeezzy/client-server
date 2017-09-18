package main

import (
	"testing"
	"net/http"
	"log"
	"encoding/json"
)

type Users struct {
	Id        string `json:id`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

const DEFAULT_HOST = "http://localhost:6060"

func TestgetUserById(t *testing.T){



}
