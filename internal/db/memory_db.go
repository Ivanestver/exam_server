package db

import (
	"exam_server/internal/structs"
	"time"
)

type _MemoryDB struct {
	messages []*structs.Message
	users    []*structs.User
}

func (memoryDB *_MemoryDB) GetAllMessages() []*structs.Message {
	messages := make([]*structs.Message, len(memoryDB.messages))
	for i, m := range memoryDB.messages {
		messages[i] = &structs.Message{
			IncomingMessage: structs.IncomingMessage{
				Owner: m.Owner,
				Text:  m.Text,
			},
			ReceivingTime: m.ReceivingTime,
		}
	}
	return messages
}

func (memoryDB *_MemoryDB) GetAllMessagesAfterTime(time time.Time) []*structs.Message {
	messages := make([]*structs.Message, 0)
	for _, m := range memoryDB.messages {
		if m.GetTimeAsTime().Compare(time) > 0 {
			messages = append(messages, &structs.Message{
				IncomingMessage: structs.IncomingMessage{
					Owner: m.Owner,
					Text:  m.Text,
				},
				ReceivingTime: m.ReceivingTime,
			})
		}
	}
	return messages
}

func (memoryDB *_MemoryDB) AddMessage(message *structs.Message) bool {
	memoryDB.messages = append(memoryDB.messages, &structs.Message{
		IncomingMessage: structs.IncomingMessage{
			Owner: message.Owner,
			Text:  message.Text,
		},
		ReceivingTime: message.ReceivingTime,
	})
	return true
}

func (memoryDB *_MemoryDB) AddUser(newUser *structs.User) bool {
	if memoryDB.CheckUser(newUser.Login) {
		return false
	}
	memoryDB.users = append(memoryDB.users, &structs.User{
		Name:     newUser.Name,
		Login:    newUser.Login,
		Password: newUser.Password,
	})
	return true
}

func (memoryDB *_MemoryDB) CheckUser(login string) bool {
	for _, u := range memoryDB.users {
		if u.Login == login {
			return true
		}
	}
	return false
}
