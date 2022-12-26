package repository

import (
	"context"
	"go.uber.org/zap"
	"si-test/pkg/common"
	"si-test/services/destination_product/domain/model"
)

type DestinationProductRepository struct {
	Super  common.BaseCapsule
	Logger *zap.Logger
}

type IDestinationProductRepository interface {
	UpdateDestinationProduct(product model.Product) error
}

func (d DestinationProductRepository) UpdateDestinationProduct(product model.Product) error {
	sql := `
		update destination_product
		set qty = $2,
		    selling_price = $3,
		    promo_price = $4,
		    updated_at = now()
		where id = $1;
	`
	_, err := d.Super.Database.Exec(context.Background(), sql,
		product.GetId(),
		product.GetQty(),
		product.GetSellingPrice(),
		product.GetPromoPrice(),
	)
	if err != nil {
		d.Logger.Error("error update to destination product", zap.Error(err))
		return err
	}
	return nil
}

func NewDestinationProductRepository(super common.BaseCapsule, logger *zap.Logger) IDestinationProductRepository {
	return &DestinationProductRepository{
		Super:  super,
		Logger: logger,
	}
}
