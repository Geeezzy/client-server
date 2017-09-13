package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	cli "github.com/jawher/mow.cli"
	//"github.com/jawher/mow.cli"
)

const DEFAULT_HOST = "http://localhost:8080"

type User struct {
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	app := cli.App("client-server", "Client-server on Golang")

	app.Command("get", "Run a command request for full users ", func(cmd *cli.Cmd) {
		cmd.Command("users", " get all users", cli.ActionCommand(func() {
			res, _ := http.Get(DEFAULT_HOST + "/user")

			body, _ := ioutil.ReadAll(res.Body)

			println(string(body))
		}))

		cmd.Command("user", "get user by id", cli.ActionCommand(func() {
			res, _ := http.Get(DEFAULT_HOST + "/user/" + "2") //localhost:8080/user/2
			body, _ := ioutil.ReadAll(res.Body)
			println(string(body))
		}))

	})

	app.Command("delete", "Delete user", func(cmd *cli.Cmd) {
		cmd.Command("user", "user by Id", cli.ActionCommand(func() {

			client := &http.Client{}
			req, err := http.NewRequest("DELETE", DEFAULT_HOST+"/user/"+"7", nil)
			if err != nil {
				log.Panic(err)
			}
			_, err = client.Do(req)
			if err != nil {
				log.Panic(err)
			}
		}))
	})

	app.Command("create", "Create users and ..", func(cmd *cli.Cmd) {
		cmd.Command("user", "create user", cli.ActionCommand(func() {
			u := User{
				Name:      "kleva",
				FirstName: "Kirill",
				LastName:  "Levin",
			}
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(u)
			res, _ := http.Post(DEFAULT_HOST+"/user", "application/json; charset=utf-8", b)
			io.Copy(os.Stdout, res.Body)
		}))
	})

	app.Run(os.Args)
}
