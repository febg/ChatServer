package message

type SentMessage struct {
	ID         string
	ChatID     string
	SenderID   string
	ReceiverID string
	Message    string
	TimeSent   int64
	Saved      bool
}

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
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
