package api

import (
	"log"

	"github.com/febg/ChatServer/chatroom"
	"github.com/febg/ChatServer/datastore"
)

type Control struct {
	Config ControlConfig
	Rooms  []*chatroom.ChatRoom
	DB     datastore.Datastore
}

type ControlConfig struct {
	LocalDB bool
}

func NewControl(config ControlConfig) (*Control, error) {
	c := Control{
		Config: config,
		Rooms:  []*chatroom.ChatRoom{},
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
