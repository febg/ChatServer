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

func (c *Control) HandleSavedMessage(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	msgID := v["message_id"]

	if msgID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ERROR: SaveMessage: Message ID information not complete")
		return
	}

	if !c.DB.SaveMessage(msgID) {
		log.Println("Unable to locate message for message id")
		return
	}
	print()
}

func (c *Control) HandleGetAllMessages(w http.ResponseWriter, r *http.Request) {
	db := c.DB.GetAllMessages()
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
	wg.Wait()
}

func (c *Control) HandleGetUserSentMessages(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	sID := v["user_id"]
	if sID == "" {
		log.Println("-> [ERROR] GetUserSentMessages: Unable to get user ID")
		return
	}
	sm := c.DB.GetUserSentMessages(sID)
	for _, v := range sm {
		fmt.Fprintf(w, "%+v\n\n", v)
	}
}

func (c *Control) HandleGetUserRecievedMessages(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	uID := v["user_id"]
	if uID == "" {
		log.Println("-> [ERROR] GetUserRecievedMessages: Unable to get user ID")
	}
	rm := c.DB.GetUserRecievedMessages(uID)

	for _, v := range rm {
		fmt.Fprintf(w, "%+v\n\n", v)
	}
}

func (c *Control) HandleGetUserMessages(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	uID := v["user_id"]
	if uID == "" {
		log.Println("-> [ERROR] GetUserMessages: Unable to get sender ID")
		return
	}

	um := c.DB.GetUserMessages(uID)
	for _, v := range um {
		fmt.Fprintf(w, "%+v\n\n", v)
	}

}

func (c *Control) HandleGetChatMessages(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	cID := v["chat_id"]
	if cID == "" {
		log.Println("-> [ERROR] GetChat: Unable to get chat ID")
		return
	}

}
