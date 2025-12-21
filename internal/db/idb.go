package db

import (
	"exam_server/internal/structs"
	"time"
)

type IDB interface {
	// Functions to deal with messages
	GetAllMessages() []*structs.Message
	GetAllMessagesAfterTime(time time.Time) []*structs.Message
	AddMessage(message *structs.Message) bool

	// Functions to deal with users
	AddUser(newUser *structs.User) bool
	CheckUser(login string) bool
}

var idb IDB

func GetDB() IDB {
	return idb
}

func InitDB() {
	idb = &_MemoryDB{}
}
