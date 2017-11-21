package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/febg/ChatServer/message"
)

func (c *Control) HandleConnections(w http.ResponseWriter, r *http.Request) {
	//Upgrade request to websocket
	ws, err := c.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Close connection at the end of function
	defer ws.Close()
	c.Rooms.Clients[ws] = true
	fmt.Println(c.Rooms.Clients)

	for {
		var msg message.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(c.Rooms.Clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		c.Rooms.Broadcaster <- msg
	}
	log.Println("Terminated Websocket")

}
