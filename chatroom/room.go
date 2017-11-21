package chatroom

import (
	"log"

	"github.com/febg/ChatServer/message"
	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	Clients     map[*websocket.Conn]bool
	Broadcaster chan message.Message
}

func NewChatRoom() *ChatRoom {
	cr := ChatRoom{

		Clients:     make(map[*websocket.Conn]bool),
		Broadcaster: make(chan message.Message),
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
