package datastore

import (
	"github.com/febg/ChatServer/message"
)

type Datastore interface {
}

type LocalDB struct {
	SentMessages     []message.SentMessage
	RecievedMessages []message.RecievedMessage
}

func NewLocalDB() (*LocalDB, error) {
	DB := LocalDB{
		SentMessages:     []message.SentMessage{},
		RecievedMessages: []message.RecievedMessage{},
	}
	return &DB, nil
}
