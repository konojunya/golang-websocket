package main

import (
	"log"
	"net/http"

	"github.com/trevex/golem"
)

type echoMessage struct {
	Msg string `json:"msg"`
}

func main() {
	http.HandleFunc("/ws", createRouter().Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createRouter() *golem.Router {
	router := golem.NewRouter()
	router.On("echo", echo)
	return router
}

func echo(conn *golem.Connection, data *echoMessage) {
	conn.Emit("echo", data)
}
