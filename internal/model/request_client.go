package model

type RequestClient struct {
	ID      int64  `json:"id,omitempty"` // omitempty — чтобы не включать в JSON при отправке
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Comment string `json:"comment"`
}
