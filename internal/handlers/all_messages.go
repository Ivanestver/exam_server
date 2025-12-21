package handlers

import (
	"encoding/json"
	"exam_server/internal/db"
	"exam_server/internal/structs"
	"net/http"
)

type _AllMessagesResponse struct {
	Messages []structs.Message `json:"messages"`
}

func AllMessagesHandler(writer http.ResponseWriter, req *http.Request) {
	const PARAMETER_NAME = "login"
	if !req.URL.Query().Has(PARAMETER_NAME) {
		sendResponse(http.StatusBadRequest, "No user name provided", writer)
		return
	}
	login := req.URL.Query().Get(PARAMETER_NAME)
	db := db.GetDB()
	if db == nil {
		sendResponse(http.StatusBadRequest, "No Database present", writer)
		return
	}

	if true || db.CheckUser(login) {
		allMessages := db.GetAllMessages()
		var response _AllMessagesResponse
		response.Messages = make([]structs.Message, len(allMessages))
		for i, msg := range allMessages {
			response.Messages[i] = *msg
		}
		if err := json.NewEncoder(writer).Encode(&allMessages); err != nil {
			sendResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), writer)
		}
	} else {
		sendResponse(http.StatusNotFound, "Could not find such a user", writer)
	}
}
