package model

import "si-test/pkg/exception"

const ProductDomainName = "Product"

type Product struct {
	id           int
	productName  string
	qty          int
	sellingPrice int
	promoPrice   int
}

func NewProduct(Id int, ProductName string, Qty int, SellingPrice int, PromoPrice int) (Product, error) {
	if Id <= 0 {
		return Product{}, exception.NewDomainError("invalid id", ProductDomainName)
	}

	if ProductName == "" {
		return Product{}, exception.NewDomainError("product name cannot empty", ProductDomainName)
	}

	if Qty < 0 {
		return Product{}, exception.NewDomainError("qty cannot less than 0", ProductDomainName)
	}

	if SellingPrice < 0 {
		return Product{}, exception.NewDomainError("selling price cannot less than 0", ProductDomainName)
	}

	if PromoPrice < 0 {
		return Product{}, exception.NewDomainError("promo price cannot less than 0", ProductDomainName)
	}

	return Product{
		id:           Id,
		productName:  ProductName,
		qty:          Qty,
		sellingPrice: SellingPrice,
		promoPrice:   PromoPrice,
	}, nil
}

func (p Product) GetId() int {
	return p.id
}

func (p Product) GetProductName() string {
	return p.productName
}

func (p Product) GetQty() int {
	return p.qty
}

func (p Product) GetSellingPrice() int {
	return p.sellingPrice
}

func (p Product) GetPromoPrice() int {
	return p.promoPrice
}
