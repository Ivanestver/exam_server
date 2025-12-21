package handlers

import (
	"encoding/json"
	"exam_server/internal/db"
	"exam_server/internal/structs"
	"net/http"
)

func SignUpHandler(writer http.ResponseWriter, req *http.Request) {
	if req == nil || req.Body == nil {
		// writer.WriteHeader(http.StatusBadRequest)
		// writer.Write([]byte(err.Error()))
		sendResponse(http.StatusBadRequest, "No data", writer)
		return
	}

	var user structs.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		// writer.WriteHeader(http.StatusBadRequest)
		// writer.Write([]byte(err.Error()))
		sendResponse(http.StatusBadRequest, err.Error(), writer)
		return
	}

	db := db.GetDB()
	if db == nil {
		// writer.WriteHeader(http.StatusInternalServerError)
		// writer.Write([]byte("Internal server error"))
		sendResponse(http.StatusBadRequest, "Internal server error", writer)
		return
	}

	if db.CheckUser(user.Login) {
		// writer.WriteHeader(http.StatusOK)
		// writer.Write([]byte("Log in successful"))
		sendResponse(http.StatusBadRequest, "Such a user exists", writer)
	} else {
		// writer.WriteHeader(http.StatusNotFound)
		// writer.Write([]byte("Such a user doesn't exist"))
		if db.AddUser(&user) {
			sendResponse(http.StatusCreated, "Sign up successful", writer)
		} else {
			sendResponse(http.StatusInternalServerError, "Couldn't sign up", writer)
		}
	}
}
