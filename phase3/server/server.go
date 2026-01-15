package server

import (
	"log"
	"net/http"

	"github.com/TheAditya-10/go/phase3/handlers"
	"github.com/TheAditya-10/go/phase3/middleware"
)

func StartServer() {
	// 1. Create a ServeMux (router)
	mux := http.NewServeMux()

	// 2. Register routes
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/user", handlers.GetUser)

	// 3. Wrap router with middleware
	handler := middleware.Logger(mux)

	// 4. Start HTTP server
	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}
}
