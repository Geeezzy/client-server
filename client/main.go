package main

import (
	"io/ioutil"
	"net/http"
	"os"

	cli "github.com/jawher/mow.cli"
	//"github.com/jawher/mow.cli"
)

const DEFAULT_HOST = "http://localhost:8080"

func main() {

	app := cli.App("docker", "A self-sufficient runtime for linux containers")

	app.Command("getusers", "Run a command request for full users ", func(cmd *cli.Cmd) {

		cmd.Action = func() {

			res, _ := http.Get(DEFAULT_HOST)

			body, _ := ioutil.ReadAll(res.Body)

			println(string(body))

		}

	})

	app.Run(os.Args)
}
