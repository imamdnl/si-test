package usecase

import (
	"go.uber.org/zap"
	"si-test/pkg/config"
	source_product "si-test/services/source_product/repository"
)

type SourceProductUseCase struct {
	repo source_product.ISourceProductRepository
}

type ISourceProductUseCase interface {
	GetAllProduct() ([]source_product.SourceProductDTO, error)
}

func (sp SourceProductUseCase) GetAllProduct() ([]source_product.SourceProductDTO, error) {
	data, err := sp.repo.GetAllData()
	if err != nil {
		config.Logger().Error("error get all product in usecase", zap.Error(err))
		return nil, err
	}
	return data, nil
}

func NewSourceProductUseCase(repo source_product.ISourceProductRepository) SourceProductUseCase {
	return SourceProductUseCase{repo: repo}
}
