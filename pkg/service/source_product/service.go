package source_product

type sourceProductService struct {
}

func NewSourceProductService() ServiceInterface {
	return &sourceProductService{}
}
