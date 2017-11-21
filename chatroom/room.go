package chatroom

import (
	"log"
	"net/http"

	"github.com/febg/ChatServer/message"
	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	Upgrader    websocket.Upgrader
	Clients     map[*websocket.Conn]bool
	Broadcaster chan message.SentMessage
}

func NewChatRoom() *ChatRoom {
	cr := ChatRoom{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Clients:     make(map[*websocket.Conn]bool),
		Broadcaster: make(chan message.SentMessage),
	}
	return &cr
}

func (c *ChatRoom) HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-c.Broadcaster
		// Send it out to every client that is currently connected
		for client := range c.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(c.Clients, client)
			}
		}
	}
}
