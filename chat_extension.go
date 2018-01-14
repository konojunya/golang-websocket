package main

import (
	"log"
	"net/http"

	"github.com/trevex/golem"
)

var room_manager = golem.NewRoomManager()

func main() {
	http.HandleFunc("/ws", createRouter().Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createRouter() *golem.Router {
	router := golem.NewRouter()
	router.SetConnectionExtension(NewConnection)
	router.On("join", join)
	router.On("say", say)
	return router
}

type Connection struct {
	*golem.Connection
	Name string
}

func NewConnection(conn *golem.Connection) *Connection {
	return &Connection{Connection: conn}
}

func join(conn *Connection, data *joinRequest) {
	conn.Name = data.Name
	room_manager.Join(data.Name, conn.Connection)
}

type joinRequest struct {
	Name string `json:"name"`
}

func say(conn *Connection, data *sayRequest) {
	room_manager.Emit(conn.Name, "say", &sayResponse{Msg: data.Msg})
}

type sayRequest struct {
	Msg string `json:"msg"`
}

type sayResponse struct {
	Msg string `json:"msg"`
}
