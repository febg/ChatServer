package message

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Message struct {
	ID         string
	SenderID   string
	ReceiverID string
	Message    string
	Info       Information
}

type Information struct {
	TimeSent     int64
	TimeRecieved int64
	Opened       bool
	Saved        bool
}

func NewMessage(sender string, reciever string, message string) (*Message, error) {
	mess := Message{
		ID:         uuid.NewV4().String(),
		SenderID:   sender,
		ReceiverID: reciever,
		Message:    message,
		Info: Information{
			TimeSent: time.Now().Unix(),

			Opened: false,
			Saved:  false,
		},
	}

	return &mess, nil
}
