package domain

type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Image       string  `json:"image" db:"image"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
	Category    string  `json:"category" db:"category"`
}
