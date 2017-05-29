package server

import (
	"errors"
	context "golang.org/x/net/context"

	"github.com/eggsbenjamin/groceries-tesco/domain"
	pb "github.com/eggsbenjamin/groceries-tesco/grpc"
)

type GetProductsHandler struct {
	productGetter domain.ProductGetter
}

//	constructor
func NewGetProductsHandler(productGetter domain.ProductGetter) *GetProductsHandler {
	return &GetProductsHandler{
		productGetter: productGetter,
	}
}

func (p *GetProductsHandler) GetProducts(ctx context.Context, in *pb.GetProductsRequest) (*pb.ProductsResponse, error) {
	prods, err := p.productGetter.Get(in.Barcode)
	if err != nil {
		return nil, errors.New("error retrieving products")
	}
	rsp := &pb.ProductsResponse{}
	for _, pr := range prods {
		rsp.Products = append(rsp.Products, mapProduct(pr))
	}
	return rsp, nil
}

//	maps domain product to protobuf product
func mapProduct(dp *domain.Product) *pb.Product {
	return &pb.Product{
		Gtin:        dp.GTIN,
		Description: dp.Description,
		Contents: &pb.ProductContents{
			Quantity:    float32(dp.Contents.Quantity),
			QuantityUom: dp.Contents.QuantityUOM,
			AvgMeasure:  dp.Contents.AvgMeasure,
			NetContents: dp.Contents.NetContents,
		},
		Characterisitics: &pb.ProductCharacterisitics{
			IsFood:  dp.ProductCharacteristics.IsFood,
			IsDrink: dp.ProductCharacteristics.IsDrink,
		},
	}
}
