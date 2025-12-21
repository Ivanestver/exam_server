package handlers

import (
	"encoding/json"
	"exam_server/internal/db"
	"exam_server/internal/structs"
	"net/http"
	"time"
)

func SendMessage(writer http.ResponseWriter, req *http.Request) {
	receivingTime := time.Now()
	var incomingMessage structs.IncomingMessage
	if err := json.NewDecoder(req.Body).Decode(&incomingMessage); err != nil {
		sendResponse(http.StatusInternalServerError, err.Error(), writer)
		return
	}
	db := db.GetDB()
	if db == nil {
		sendResponse(http.StatusBadRequest, "No Database present", writer)
		return
	}

	message := structs.NewMessage(incomingMessage)
	message.ReceivingTime = receivingTime.Format(time.DateTime)

	if db.CheckUser(message.Owner) {
		db.AddMessage(message)
		sendResponse(http.StatusOK, "Message has sent successfully!", writer)
	} else {
		sendResponse(http.StatusBadRequest, "You're not signed in", writer)
	}
}
