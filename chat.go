package main

import (
	"log"
	"net/http"

	"github.com/trevex/golem"
)

var room_manager = golem.NewRoomManager()

type joinRequest struct {
	Name string `json:"name"`
}

type sayRequest struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

type sayResponse struct {
	Msg string `json:"msg"`
}

func main() {
	http.HandleFunc("/ws", createRouter().Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createRouter() *golem.Router {
	router := golem.NewRouter()
	router.On("join", join)
	router.On("say", say)
	return router
}

func join(conn *golem.Connection, data *joinRequest) {
	room_manager.Join(data.Name, conn)
}

func say(conn *golem.Connection, data *sayRequest) {
	room_manager.Emit(data.Name, "say", &sayResponse{Msg: data.Msg})
}
