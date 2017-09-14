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
		//СДЕЛАТЬ ФЛАГ ID
		cmd.Command("user", "get user by id", func(sc *cli.Cmd) {
			//доописать ввод id с клавы
			sc.Spec = "[-r] DST "

			var (
				recursive = sc.BoolOpt("r recursive", false, "Copy files recursively")
				src       = sc.StringArg("DST", "", "Destination where to copy files to")
			)
			sc.Action = func() {
				fmt.Printf("Copying %s to [recursively: %v ]\n", *src, *recursive)
				res, err := http.Get(DEFAULT_HOST + "/user/" + *src)
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

	//СДЕЛАТЬ ФЛАГ ID

	app.Command("delete", "Delete user", func(cmd *cli.Cmd) {
		cmd.Command("user", "user by Id", func(sc *cli.Cmd) {
			//ввод с клавы
			sc.Spec = "[-r] DST "

			var (
				recursive = sc.BoolOpt("r recursive", false, "Copy files recursively")
				src       = sc.StringArg("DST", "", "Destination where to copy files to")
			)
			sc.Action = func() {
				fmt.Printf("Copying %s to [recursively: %v ]\n", *src, *recursive)
				client := &http.Client{}
				req, err := http.NewRequest("DELETE", DEFAULT_HOST+"/user/"+*src, nil)
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
		//переписать с files.go
		cmd.Command("user", "create user", func(sc *cli.Cmd) {

			sc.Spec = "[-r] DST "

			var (
				recursive = sc.BoolOpt("r recursive", false, "Copy files recursively")
				src       = sc.StringArg("DST", "", "Destination where to copy files to")
			)

			sc.Action = func() {
				fmt.Printf("Copying %s to [recursively: %v ]\n", *src, *recursive)

				bs, err := ioutil.ReadFile(*src)
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

	//СДЕЛАТЬ ФЛАГ ID и ПУТЬ ДО ФАЙЛА

	app.Command("update", "Update  ", func(cmd *cli.Cmd) {
		//переписать с files.go и ввод id
		cmd.Command("user", "Update user", func(sc *cli.Cmd) {
			sc.Spec = "[-r] DST "
			var (
				recursive = sc.BoolOpt("r recursive", false, "Copy files recursively")
				src       = sc.StringArg("DST", "", "Destination where to copy files to")
			)

			sc.Action = func() {
				bs, err := ioutil.ReadFile(*src)
				if err != nil {
					panic(err)
					return
				}

				b := bytes.NewBuffer(bs)
				fmt.Printf("Copying %s to [recursively: %v ]\n", *src, *recursive)

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
			}
		})
	})

	app.Run(os.Args)
}
