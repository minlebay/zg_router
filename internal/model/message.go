package model

import "time"

type Message struct {
	UUID           string
	ContentType    string
	MessageContent MessageContent
}

type MessageContent struct {
	SendAt   time.Time
	Provider string
	Consumer string
	Title    string
	Content  string
}
