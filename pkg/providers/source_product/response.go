package source_product

type DataSourceProduct struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []Data `json:"data"`
}

type Data struct {
	ID           int    `json:"id"`
	ProductName  string `json:"product_name"`
	Qty          int    `json:"qty"`
	SellingPrice int    `json:"selling_price"`
	PromoPrice   int    `json:"promo_price"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
