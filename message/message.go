package message

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SentMessage struct {
	ID         string
	ChatID     string
	SenderID   string
	ReceiverID string
	Message    string
	TimeSent   int64
	Saved      bool
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

func NewMessage(sender string, reciever string, message string) (*SentMessage, *RecievedMessage, error) {

	chatID := uuid.NewV4().String()

	sm := SentMessage{
		ID:         uuid.NewV4().String(),
		ChatID:     chatID,
		SenderID:   sender,
		ReceiverID: reciever,
		Message:    message,
		TimeSent:   time.Now().Unix(),
		Saved:      false,
	}

	rm := RecievedMessage{
		ID:           uuid.NewV4().String(),
		ChatID:       chatID,
		ReceiverID:   reciever,
		SenderID:     sender,
		Message:      message,
		TimeRecieved: time.Now().Unix(),
		Saved:        false,
		Info: Information{
			Opened: false,
		},
	}

	return &sm, &rm, nil
}
