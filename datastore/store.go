package datastore

import (
	"github.com/febg/ChatServer/message"
)

type Datastore interface {
	StoreSentMessage(message.SentMessage) error
	StoreRecievedMessage(message.RecievedMessage) error
	GetMessages() error
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

func (db *LocalDB) StoreSentMessage(sm message.SentMessage) error {
	db.SentMessages = append(db.SentMessages, sm)
	return nil

}

func (db *LocalDB) StoreRecievedMessage(sm message.RecievedMessage) error {
	return nil
}

func (db *LocalDB) GetMessages() error {

	return nil
}
