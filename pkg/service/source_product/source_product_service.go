package source_product

import (
	"si-test/pkg/providers/source_product"
)

type ServiceInterface interface {
	GetAllSourceProducts() (*source_product.DataSourceProduct, error)
}

func (sp sourceProductService) GetAllSourceProducts() (*source_product.DataSourceProduct, error) {
	respSP, err := source_product.SourceProductProvider.GetAllSourceProduct()
	if err != nil {
		return nil, err
	}

	return respSP, nil
}
