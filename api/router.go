package api

import "github.com/gorilla/mux"

func Router(c *Control) *mux.Router {
	r := mux.NewRouter()
	r.Methods("GET").Path("/ws").HandlerFunc(c.HandleConnections)
	r.Methods("GET").Path("/saveMessage/{}").HandlerFunc(c.HandleSavedMessages)
	r.Methods("GET").Path("/getAllMessages/").HandlerFunc(c.HandleGetMessages)
	return r
}
