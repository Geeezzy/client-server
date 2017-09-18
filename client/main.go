package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	cli "github.com/jawher/mow.cli"

)

const DEFAULT_HOST = "http://localhost:6060"

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

		cmd.Command("user", "get user by id", func(sc *cli.Cmd) {
			//доописать ввод id с клавы
			sc.Spec = "[-i] ID "

			var (
				checkId = sc.BoolOpt("i id", false, "Read id")
				id       = sc.StringArg("ID", "", "What id to use")
			)
			sc.Action = func() {
				fmt.Printf("get id %s to [sucsessful: %v ]\n", *id, *checkId)
				res, err := http.Get(DEFAULT_HOST + "/user/" + *id)
				if err != nil {
					log.Panic(err)
				}
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Panic(err)
				}
				println(string(body))
			}
		})

	})

	app.Command("delete", "Delete user", func(cmd *cli.Cmd) {
		cmd.Command("user", "user by Id", func(sc *cli.Cmd) {
			//ввод с клавы
			sc.Spec = "[-i] ID "

			var (
				checkId = sc.BoolOpt("i id", false, "Read id")
				id       = sc.StringArg("ID", "", "What id to use")
			)
			sc.Action = func() {
				fmt.Printf("Delete id %s to [sucsessful: %v ]\n", *id, *checkId)
				client := &http.Client{}
				req, err := http.NewRequest("DELETE", DEFAULT_HOST + "/user/"+*id, nil)
				if err != nil {
					log.Panic(err)
				}
				_, err = client.Do(req)
				if err != nil {
					log.Panic(err)
				}
			}
		})
	})

	app.Command("create", "Create users and ..", func(cmd *cli.Cmd) {
		cmd.Command("user", "create user", func(sc *cli.Cmd) {

			sc.Spec = "[-f] PATH "

			var (
				checkPath = sc.BoolOpt("f force", false, "Read path")
				path       = sc.StringArg("PATH", "", "The path to the file")
			)

			sc.Action = func() {
				fmt.Printf("Create user to [sucsessful: %v ]\n",  *checkPath)

				bs, err := ioutil.ReadFile(*path)
				if err != nil {
					panic(err)
					return
				}

				b := bytes.NewBuffer(bs)

				res, err := http.Post(DEFAULT_HOST+"/user", "application/json; charset=utf-8", b)
				if err != nil {
					log.Panic(err)
				}
				io.Copy(os.Stdout, res.Body)
			}
		})
	})

	app.Command("update", "Update  ", func(cmd *cli.Cmd) {
		cmd.Command("user", "Update user", func(sc *cli.Cmd) {
			sc.Spec = "[-f] ID PATH "
			var (
				checkId = sc.BoolOpt("f force", false, "Read id and path")
				id 	= sc.StringArg("ID", "", "What id to use")
				path       = sc.StringArg("PATH", "", "The path to the file")
			)

			sc.Action = func() {
				bs, err := ioutil.ReadFile(*path)
				if err != nil {
					panic(err)
					return
				}

				b := bytes.NewBuffer(bs)
				fmt.Printf("Update %s to [sucsessful: %v ] , %s\n", *id, *checkId, *path)

				client := &http.Client{}
				req, err := http.NewRequest("PUT", DEFAULT_HOST+"/user/"+*id, b)
				if err != nil {
					log.Panic(err)
				}
				_, err = client.Do(req)
				if err != nil {
					log.Panic(err)
				}
				
			}
		})
	})

	app.Run(os.Args)
}
