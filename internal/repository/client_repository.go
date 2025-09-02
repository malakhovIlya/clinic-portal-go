package repository

import (
	"github.com/malakhovIlya/clinic-portal-go/internal/model"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) Save(request model.RequestClient) error {
	return r.db.Create(&request).Error
}
