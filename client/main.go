package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
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

		cmd.Command("user id", "het user by id", cli.ActionCommand(func() {
			res, _ := http.Get(DEFAULT_HOST + "/user" + "/2")
			body, _ := ioutil.ReadAll(res.Body)
			println(string(body))
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
