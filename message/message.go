package message

import "time"

type Message interface {
	SetCurrentTime()
}

// {"message_id":"","chat_id":"","sender_id":"","reciever_id":"message":"","time_saved":"","message":""}

type SentMessage struct {
	ID         string `json:"message_id"`
	ChatID     string `json:"chat_id"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"reciever_id"`
	Message    string `json:"message_id"`
	TimeSent   int64  `json:"time_sent"`
	Saved      bool   `json:"saved"`
}

type RecievedMessage struct {
	ID           string
	ChatID       string
	ReceiverID   string
	SenderID     string
	Message      string
	TimeRecieved int64
	Saved        bool
	Info         Information
}

type Information struct {
	Opened     bool
	TimeOpened int64
}

func (sm SentMessage) SetCurrentTime() {
	sm.TimeSent = time.Now().Unix()
}

func (rm RecievedMessage) SetCurrentTime() {
	rm.TimeRecieved = time.Now().Unix()
}
