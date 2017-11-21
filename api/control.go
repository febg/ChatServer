package api

import (
	"log"
	"net/http"

	"github.com/febg/ChatServer/chatroom"
	"github.com/febg/ChatServer/datastore"
	"github.com/gorilla/websocket"
)

type Control struct {
	Config   ControlConfig
	Rooms    *chatroom.ChatRoom
	DB       datastore.Datastore
	Upgrader websocket.Upgrader
}

type ControlConfig struct {
	LocalDB bool
}

func NewControl(config ControlConfig) (*Control, error) {
	c := Control{
		Config: config,

		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	var err error
	if config.LocalDB {
		c.DB, err = datastore.NewLocalDB()
		if err != nil {
			log.Printf("-> [FATAL] Could not get local database: %v", err)
			return nil, nil
		}
	}

	return &c, nil
}
