package model

import "time"

type RequestClient struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Name      string    `json:"name"`
	Phone     string    `json:"phoneNumber"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
}
