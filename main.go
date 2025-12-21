package main

import (
	"exam_server/internal/db"
	"exam_server/internal/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/all-messages", handlers.AllMessagesHandler)
	router.Post("/all-messages", handlers.AllMessagesAferTime)
	router.Post("/sign-in", handlers.SignInHandler)
	router.Post("/sign-up", handlers.SignUpHandler)
	router.Post("/send-message", handlers.SendMessage)

	db.InitDB()

	fmt.Println("Server started listening on port 50000")
	http.ListenAndServe(":50000", router)
}
