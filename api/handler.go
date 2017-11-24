package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/febg/ChatServer/message"
	"github.com/gorilla/mux"
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

		msg.SetCurrentTime()

		if msg.ReceiverID == "" || msg.Message == "" || msg.SenderID == "" {
			log.Printf("-> [ERROR] Message information not complete %+v", msg)
			break
		}

		// Store SentMessage and RecievedMessage
		c.DB.StoreSentMessage(msg)
		c.DB.StoreRecievedMessage(msg)
		// Send the newly received message to the broadcast channel

		c.Rooms.Broadcaster <- msg
	}
	log.Println(c.DB)

}

func (c *Control) HandleSavedMessages(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	msgID := v["message_id"]

	if msgID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ERROR: Message ID information not complete")
		return
	}
}

func (c *Control) HandleGetMessages(w http.ResponseWriter, r *http.Request) {
	db := c.DB.GetSentMessages()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range db.SentMessages {
			fmt.Fprintf(w, "%+v\n\n", v)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range db.RecievedMessages {
			fmt.Fprintf(w, "%+v\n\n", v)
		}
	}()

}
