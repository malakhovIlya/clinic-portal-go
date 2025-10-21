package api

import (
	"encoding/json"
	"github.com/malakhovIlya/clinic-portal-go/internal/model"
	"github.com/malakhovIlya/clinic-portal-go/internal/repository"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type ClientHandler struct {
	repo *repository.ClientRepository
}

func NewClientHandler(db *gorm.DB) *ClientHandler {
	return &ClientHandler{repo: repository.NewClientRepository(db)}
}

func (h *ClientHandler) SaveClientRequestHandler(w http.ResponseWriter, r *http.Request) {
	var req model.RequestClient

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Ошибка парсинга JSON: %v", err)
		http.Error(w, "Произошла ошибка!", 418)
		return
	}

	if err := h.repo.Save(req); err != nil {
		log.Printf("Ошибка сохранения в БД: %v", err)
		http.Error(w, "Произошла ошибка!", 418)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Заявка успешно создана!"))
}
