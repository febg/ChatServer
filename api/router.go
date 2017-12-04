package api

import "github.com/gorilla/mux"

func Router(c *Control) *mux.Router {
	r := mux.NewRouter()
	r.Methods("GET").Path("/ws").HandlerFunc(c.HandleConnections)
	r.Methods("GET").Path("/saveMessage/{message_id}").HandlerFunc(c.HandleSavedMessage)
	r.Methods("GET").Path("/getAllMessages/").HandlerFunc(c.HandleGetAllMessages)
	r.Methods("GET").Path("/getSentMessages/{user_id}").HandlerFunc(c.HandleGetUserSentMessages)
	r.Methods("GET").Path("/getUserMessages/{user_id}").HandlerFunc(c.HandleGetUserMessages)
	r.Methods("GET").Path("/getAllUserMessages/{user_id}").HandlerFunc(c.HandleGetUserRecievedMessages)
	r.Methods("GET").Path("/getChatMessages/{chat_id}").HandlerFunc(c.HandleGetChatMessages)

	return r
}
