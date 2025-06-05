package model

type Purchase struct {
	ID        int     `json:"id" gorm:"primary_key"`
	UserID    int     `json:"-"`
	Item      string  `json:"item"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
}
