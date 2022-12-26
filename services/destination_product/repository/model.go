package repository

import (
	"si-test/services/destination_product/domain/model"
	"time"
)

type DestinationProduct struct {
	ID           int
	ProductName  string
	Qty          int
	SellingPrice int
	PromoPrice   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUpdateDestinationProduct(model model.Product) *DestinationProduct {
	return &DestinationProduct{
		ID:           model.GetId(),
		ProductName:  model.GetProductName(),
		Qty:          model.GetQty(),
		SellingPrice: model.GetSellingPrice(),
		PromoPrice:   model.GetPromoPrice(),
		UpdatedAt:    time.Now(),
	}
}
