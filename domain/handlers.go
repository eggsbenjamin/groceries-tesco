package domain

type ProductGetter interface {
	Get(barcode string) ([]*Product, error)
}

type TescoClient interface {
	GetProductsByGTIN(string) ([]*Product, error)
}

type ProductHandler struct {
	tClient TescoClient
}

func (p *ProductHandler) Get(barcode string) ([]*Product, error) {
	return p.tClient.GetProductsByGTIN(barcode)
}

// Constructor
func NewProductHandler(tClient TescoClient) *ProductHandler {
	return &ProductHandler{
		tClient: tClient,
	}
}
