package source_product

import (
	"encoding/json"
	"errors"
	"github.com/avast/retry-go"
	"os"
	"si-test/pkg/clients/source_product"
	"time"
)

const (
	apiName = "source-product"
	url     = "/source-product/get-all"
)

type sourceProductProvider struct{}

type sourceProductProviderInterface interface {
	GetAllSourceProduct() (*DataSourceProduct, error)
}

var (
	SourceProductProvider sourceProductProviderInterface = &sourceProductProvider{}
)

func (s sourceProductProvider) GetAllSourceProduct() (*DataSourceProduct, error) {
	var result *DataSourceProduct
	if os.Getenv("SP_API_CALL") != "1" {
		return &DataSourceProduct{
			Code:    "200",
			Message: "Success",
			Data: []Data{
				{
					ID:           1,
					ProductName:  "Satu",
					Qty:          20,
					SellingPrice: 12310,
					PromoPrice:   220,
					CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
					UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
				},
				{
					ID:           2,
					ProductName:  "Dua",
					Qty:          10,
					SellingPrice: 17000,
					PromoPrice:   123,
					CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
					UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
				},
			},
		}, nil
	}
	err := retry.Do(
		func() error {
			response, err := source_product.Get(apiName, url, nil)
			if err != nil {
				return errors.New(err.Error())
			}
			if err := json.Unmarshal(response, &result); err != nil {
				return errors.New("can not unmarshal JSON")
			}

			return nil
		}, retry.Attempts(3), retry.DelayType(retry.FixedDelay),
	)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	return result, nil
}
