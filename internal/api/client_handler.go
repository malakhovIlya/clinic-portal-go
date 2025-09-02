package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/malakhovIlya/clinic-portal-go/internal/model"
)

func SaveClientRequestHandler(w http.ResponseWriter, r *http.Request) {
	var req model.RequestClient

	// Парсим JSON
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Ошибка парсинга JSON: %v", err)
		http.Error(w, "Произошла ошибка!", 418)
		return
	}

	log.Printf("Принята заявка от клиента: %+v", req)

	// Пока без базы: просто эмулируем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Заявка успешно создана!"))
}
