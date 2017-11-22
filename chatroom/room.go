package chatroom

import (
	"log"

	"github.com/febg/ChatServer/message"
	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	Clients     map[string]*websocket.Conn
	Broadcaster chan message.SentMessage
}

func NewChatRoom() *ChatRoom {
	cr := ChatRoom{

		Clients:     make(map[string]*websocket.Conn),
		Broadcaster: make(chan message.SentMessage),
	}
	return &cr
}

func (c *ChatRoom) HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-c.Broadcaster
		// Send it out to every client that is currently connected
		rAdd := c.Clients[msg.ReceiverID]

		if rAdd != nil {
			err := rAdd.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				rAdd.Close()
				delete(c.Clients, msg.ReceiverID)
			}
		} else {
			log.Println("[Error] User not found")
		}
	}
}
