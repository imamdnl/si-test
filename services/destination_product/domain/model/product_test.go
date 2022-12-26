package model_test

import (
	"github.com/stretchr/testify/assert"
	"si-test/services/destination_product/domain/model"
	"testing"
)

func Test_NewProduct(t *testing.T) {
	id := 1
	productName := "Sirup"
	qty := 10
	sellingPrice := 1000
	promoPrice := 250

	product, err := model.NewProduct(id, productName, qty, sellingPrice, promoPrice)
	assert.Nil(t, err)
	assert.Equal(t, productName, product.GetProductName())
	assert.Equal(t, qty, product.GetQty())
	assert.Equal(t, sellingPrice, product.GetSellingPrice())
	assert.Equal(t, promoPrice, product.GetPromoPrice())
}

func Test_ValidateProduct(t *testing.T) {
	invalidId := 0
	productName := "Sirup"
	qty := 10
	sellingPrice := 1000
	promoPrice := 250
	type args struct {
		id           int
		productName  string
		qty          int
		sellingPrice int
		promoPrice   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "invalid id",
			args: args{
				id:           invalidId,
				productName:  productName,
				qty:          qty,
				sellingPrice: sellingPrice,
				promoPrice:   promoPrice,
			},
		},
		{
			name: "product name cannot empty",
			args: args{
				id:           1,
				productName:  "",
				qty:          qty,
				sellingPrice: sellingPrice,
				promoPrice:   promoPrice,
			},
		},
		{
			name: "qty cannot less than 0",
			args: args{
				id:           1,
				productName:  productName,
				qty:          -1,
				sellingPrice: sellingPrice,
				promoPrice:   promoPrice,
			},
		},
		{
			name: "selling price cannot less than 0",
			args: args{
				id:           1,
				productName:  productName,
				qty:          qty,
				sellingPrice: -1,
				promoPrice:   promoPrice,
			},
		},
		{
			name: "promo price cannot less than 0",
			args: args{
				id:           1,
				productName:  productName,
				qty:          qty,
				sellingPrice: sellingPrice,
				promoPrice:   -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := model.NewProduct(tt.args.id, tt.args.productName, tt.args.qty, tt.args.sellingPrice,
				tt.args.promoPrice)
			assert.NotNil(t, err)
		})
	}
}
