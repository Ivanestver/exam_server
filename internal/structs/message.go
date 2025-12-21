package structs

import "time"

type IncomingMessage struct {
	Owner string `json:"owner"`
	Text  string `json:"text"`
}

type Message struct {
	IncomingMessage
	ReceivingTime string `json:"time"`
}

func NewMessage(incomingMessage IncomingMessage) *Message {
	return &Message{
		IncomingMessage: incomingMessage,
		ReceivingTime:   time.Now().Format(time.DateTime),
	}
}

func ToTime(t string) (time.Time, error) {
	return time.Parse(time.DateTime, t)
}

func (msg *Message) GetTimeAsTime() time.Time {
	if t, e := ToTime(msg.ReceivingTime); e == nil {
		return t
	} else {
		return time.Now()
	}
}
