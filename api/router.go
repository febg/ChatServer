package api

import "github.com/gorilla/mux"

func Router(c *Control) *mux.Router {
	r := mux.NewRouter()
	r.Methods("GET").Path("/ws").HandlerFunc(c.HandleConnections)
	r.Methods("GET").Path("/saveMessage/{message_id}").HandlerFunc(c.HandleSavedMessages)
	r.Methods("GET").Path("/getAllMessages/").HandlerFunc(c.HandleGetAllMessages)
	r.Methods("GET").Path("/getSentMessages/{sender_id}").HandlerFunc(c.HandleGetSentMessages)
	r.Methods("GET").Path("/getUserMessages/{user_id}").HandlerFunc(c.HandleGetUserMessages)
	return r
}
