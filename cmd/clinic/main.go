package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/malakhovIlya/clinic-portal-go/internal/api"
)

func main() {
	router := chi.NewRouter()

	// API
	router.Post("/api/client/request/save", api.SaveClientRequestHandler)

	// Static frontend
	fs := http.FileServer(http.Dir("web/dist"))
	router.Handle("/*", fs)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
