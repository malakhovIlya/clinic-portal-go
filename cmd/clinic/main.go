package main

import (
	"github.com/malakhovIlya/clinic-portal-go/internal/config"
	"github.com/malakhovIlya/clinic-portal-go/internal/model"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/malakhovIlya/clinic-portal-go/internal/api"
)

func main() {
	db := config.InitDB()
	db.AutoMigrate(&model.RequestClient{})

	handler := api.NewClientHandler(db)

	router := chi.NewRouter()
	router.Post("/api/client/request/save", handler.SaveClientRequestHandler)

	fs := http.FileServer(http.Dir("web/dist"))
	router.Handle("/*", fs)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
