package usecase

import (
	"go.uber.org/zap"
	"si-test/pkg/config"
	source_product2 "si-test/pkg/providers/source_product"
	"si-test/pkg/service/source_product"
	"si-test/services/destination_product/domain/model"
	destination_product "si-test/services/destination_product/repository"
)

type DestinationProductUseCase struct {
	repo    destination_product.IDestinationProductRepository
	service source_product.ServiceInterface
}

type IDestinationProductUseCase interface {
	Update() error
}

func (dp DestinationProductUseCase) Update() error {
	results, err := dp.service.GetAllSourceProducts()
	if err != nil {
		config.Logger().Error("error get all source product in usecase", zap.Error(err))
		return err
	}
	sProducts := results.Data
	for _, v := range sProducts {
		go func(v source_product2.Data) {
			product, err := model.NewProduct(v.ID, v.ProductName, v.Qty, v.SellingPrice, v.PromoPrice)
			if err != nil {
				config.Logger().Error("error create domain product in usecase", zap.Error(err))
				return
			}
			err = dp.repo.UpdateDestinationProduct(product)
			if err != nil {
				config.Logger().Error("error update destination product in usecase", zap.Error(err))
				return
			}
		}(v)
	}
	return nil
}

func NewDestinationProductUseCase(
	repo destination_product.IDestinationProductRepository,
	service source_product.ServiceInterface,
) DestinationProductUseCase {
	return DestinationProductUseCase{
		repo:    repo,
		service: service,
	}
}
