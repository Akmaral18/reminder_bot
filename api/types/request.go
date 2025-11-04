package types

type Update struct {
	Message *Message `json:"message,omitempty"`
}

type Message struct {
	ID   int    `json:"message_id"`
	From *User  `json:"from,omitempty"`
	Chat *Chat  `json:"chat,omitempty"`
	Text string `json:"text,omitempty"`
}

type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type SendMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
