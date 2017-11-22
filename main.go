package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/febg/ChatServer/api"
	"github.com/febg/ChatServer/chatroom"
)

func main() {
	// Create a simple file server

	control, err := api.NewControl(api.ControlConfig{
		LocalDB: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	control.Rooms = chatroom.NewChatRoom()

	go control.Rooms.HandleMessages()

	fmt.Println(control)
	router := api.Router(control)
	log.Println("http server started on :localhost:8000")
	err = http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
