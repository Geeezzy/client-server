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
			res, err := http.Get(DEFAULT_HOST + "/user")
			if err != nil {
				log.Panic(err)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Panic(err)
			}

			println(string(body))
		}))

		cmd.Command("user", "get user by id", cli.ActionCommand(func() {
			//доописать ввод id с клавы
			res, err := http.Get(DEFAULT_HOST + "/user/" + "2")
			if err != nil {
				log.Panic(err)
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Panic(err)
			}
			println(string(body))
		}))

	})

	app.Command("delete", "Delete user", func(cmd *cli.Cmd) {
		cmd.Command("user", "user by Id", cli.ActionCommand(func() {
			//ввод с клавы
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
		//переписать с files.go
		cmd.Command("user", "create user", cli.ActionCommand(func() {

			bs, err := ioutil.ReadFile("files/name.json")
			if err != nil {
				return
			}
			b := bytes.NewBuffer(bs)

			res, err := http.Post(DEFAULT_HOST+"/user", "application/json; charset=utf-8", b)
			if err != nil {
				log.Panic(err)
			}
			io.Copy(os.Stdout, res.Body)
		}))
	})

	app.Command("update", "Update  ", func(cmd *cli.Cmd) {
		//переписать с files.go и ввод id
		cmd.Command("user", "Update user", cli.ActionCommand(func() {
			u := User{
				Name:      "lex",
				FirstName: "Petr",
				LastName:  "Moooo",
			}
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(u)

			client := &http.Client{}
			req, err := http.NewRequest("PUT", DEFAULT_HOST+"/user/"+"2", b)
			if err != nil {
				log.Panic(err)
			}
			_, err = client.Do(req)
			if err != nil {
				log.Panic(err)
			} else {
				println("Changes added")
			}
		}))
	})

	app.Run(os.Args)
}
