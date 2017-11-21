package api

import "github.com/gorilla/mux"

func Router(c *Control) *mux.Router {
	r := mux.NewRouter()
	r.Methods("GET").Path("/ws").HandlerFunc(c.HandleConnections)
	return r
}
