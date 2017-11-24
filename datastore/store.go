package datastore

import "github.com/febg/ChatServer/message"

type Datastore interface {
	StoreSentMessage(message.SentMessage) error
	StoreRecievedMessage(message.SentMessage) error
	GetAllMessages() *LocalDB
	GetSentMessages(string) []message.SentMessage
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

func (db *LocalDB) StoreRecievedMessage(sm message.SentMessage) error {
	rm := message.RecievedMessage{
		ChatID:     sm.ChatID,
		ReceiverID: sm.ReceiverID,
		SenderID:   sm.SenderID,
		Message:    sm.Message,
		Saved:      false,
		Info: message.Information{
			Opened: false,
		},
	}
	rm.SetCurrentTime()
	db.RecievedMessages = append(db.RecievedMessages, rm)
	return nil
}

func (db *LocalDB) GetAllMessages() *LocalDB {
	return db
}

func (db *LocalDB) GetSentMessages(id string) []message.SentMessage {
	sm := []message.SentMessage{}
	for _, v := range db.SentMessages {
		if v.SenderID == id {
			sm = append(sm, v)
		}
	}
	return sm
}
