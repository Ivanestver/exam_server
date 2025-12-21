package handlers

import (
	"encoding/json"
	"exam_server/internal/db"
	"exam_server/internal/structs"
	"net/http"
)

type _AllMessagesAfterTimeRequest struct {
	Login string `json:"login"`
	Time  string `json:"time"`
}

func AllMessagesAferTime(writer http.ResponseWriter, req *http.Request) {
	var request _AllMessagesAfterTimeRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		sendResponse(http.StatusInternalServerError, err.Error(), writer)
		return
	}
	login := request.Login
	timeAfter, err := structs.ToTime(request.Time)
	if err != nil {
		sendResponse(http.StatusBadRequest, err.Error(), writer)
		return
	}
	db := db.GetDB()
	if db == nil {
		sendResponse(http.StatusBadRequest, "No Database present", writer)
		return
	}

	if true || db.CheckUser(login) {
		allMessages := db.GetAllMessagesAfterTime(timeAfter)
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
