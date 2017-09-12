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

	app.Command("getusers", "Run a command request for full users ", func(cmd *cli.Cmd) {

		cmd.Action = func() {

			res, _ := http.Get(DEFAULT_HOST + "/getallusers")

			body, _ := ioutil.ReadAll(res.Body)

			println(string(body))

		}

	})

	app.Command("create", "Create users and ..", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			u := User{
				Name:      "test",
				FirstName: "lol",
				LastName:  "kek",
			}
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(u)
			res, _ := http.Post(DEFAULT_HOST+"/createuser", "application/json; charset=utf-8", b)
			io.Copy(os.Stdout, res.Body)
		}
	})

	app.Run(os.Args)
}
