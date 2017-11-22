package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/febg/ChatServer/message"
	uuid "github.com/satori/go.uuid"
)

func (c *Control) HandleConnections(w http.ResponseWriter, r *http.Request) {
	//Upgrade request to websocket
	ws, err := c.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Close connection at the end of function
	defer ws.Close()
	connectionID := uuid.NewV4().String()
	c.Rooms.Clients[connectionID] = ws
	fmt.Println(c.Rooms.Clients)

	for {
		var msg message.SentMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(c.Rooms.Clients, connectionID)
			break
		}
		msg.TimeSent = time.Now().Unix()
		log.Println(msg)
		// Send the newly received message to the broadcast channel
		c.Rooms.Broadcaster <- msg
	}
	log.Println("Terminated Websocket")

}
