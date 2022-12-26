package source_product

import (
	"time"
)

type SourceProductDTO struct {
	ID           int       `json:"id"`
	ProductName  string    `json:"product_name"`
	Qty          int       `json:"qty"`
	SellingPrice int       `json:"selling_price"`
	PromoPrice   int       `json:"promo_price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
