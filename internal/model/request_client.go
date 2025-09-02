package model

type RequestClient struct {
	ID    int64  `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
