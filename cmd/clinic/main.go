package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/malakhovIlya/clinic-portal-go/internal/config"
	"github.com/malakhovIlya/clinic-portal-go/internal/model"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/malakhovIlya/clinic-portal-go/internal/api"
)

const (
	distDir = "web/dist"
)

func main() {
	db := config.InitDB()
	db.AutoMigrate(&model.RequestClient{})

	handler := api.NewClientHandler(db)

	router := chi.NewRouter()

	// Логирование запросов
	router.Use(middleware.Logger)

	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://127.0.0.1:8080", "http://your-domain.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // кэш preflight
	}))

	router.Post("/api/client/request/save", handler.SaveClientRequestHandler)

	fs := http.FileServer(http.Dir("web/dist"))
	router.Handle("/*", fs)

	// --- WEB ---
	attachWebRoutes(router)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func attachWebRoutes(router *chi.Mux) {
	// 1) Раздаём файлы из dist (собранный Vue)
	fs := http.FileServer(http.Dir(distDir))
	router.Handle("/*", fs)

	// 2) Любой путь без точки → кастомный 404
	router.Get("/{path:[^.]*}", func(w http.ResponseWriter, r *http.Request) {
		serveCustom404(w, r)
	})
}

func serveCustom404(w http.ResponseWriter, r *http.Request) {
	// Если в URL есть точка (например, /main.js), то это не 404, пусть обслужит FileServer
	if strings.Contains(filepath.Base(r.URL.Path), ".") {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	http.ServeFile(w, r, filepath.Join(distDir, "404.html"))
}
