package datastore

import "github.com/febg/ChatServer/message"

type Datastore interface {
	StoreSentMessage(message.SentMessage) error
	StoreRecievedMessage(message.SentMessage) error
	GetAllMessages() *LocalDB
	GetUserSentMessages(string) []message.SentMessage
	GetUserRecievedMessages(string) []message.RecievedMessage
	GetUserMessages(string) []message.Message
	SaveMessage(string) bool
	GetChatMessages(string) []message.Message
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

func (db *LocalDB) GetUserSentMessages(id string) []message.SentMessage {
	sm := []message.SentMessage{}
	for _, v := range db.SentMessages {
		if v.SenderID == id {
			sm = append(sm, v)
		}
	}
	return sm
}

func (db *LocalDB) GetUserRecievedMessages(id string) []message.RecievedMessage {
	rm := []message.RecievedMessage{}
	for _, v := range db.RecievedMessages {
		if v.ID == id {
			rm = append(rm, v)
		}
	}
	return rm
}

func (db *LocalDB) GetUserMessages(id string) []message.Message {
	m := []message.Message{}
	for _, v := range db.SentMessages {
		if v.SenderID == id {
			m = append(m, v)
		}
	}

	for _, v := range db.RecievedMessages {
		if v.ReceiverID == id {
			m = append(m, v)
		}

	}
	return m
}

func (db *LocalDB) SaveMessage(mID string) bool {
	for _, v := range db.SentMessages {
		if v.ID == mID {
			sa := &v
			sa.Saved = true
			return true
		}
	}
	for _, v := range db.RecievedMessages {
		if v.ID == mID {
			sa := &v
			sa.Saved = true
			return true
		}
	}
	return false
}

func (db *LocalDB) GetChatMessages(cID string) []message.Message {
	cm := []message.Message{}
	for _, v := range db.SentMessages {
		if v.ChatID == cID {
			cm = append(cm, v)
		}
	}
	for _, v := range db.RecievedMessages {
		if v.ChatID == cID {
			cm = append(cm, v)
		}
	}
	return cm
}
