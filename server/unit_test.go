// +build unit

package server_test

import (
	"errors"

	"github.com/eggsbenjamin/groceries-tesco/domain"
	"github.com/eggsbenjamin/groceries-tesco/domain/mocks"
	pb "github.com/eggsbenjamin/groceries-tesco/grpc"
	. "github.com/eggsbenjamin/groceries-tesco/server"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	It("should return the correct product response", func() {
		By("setup")
		mockGetter := &mocks.ProductGetter{}
		mockProducts := []*domain.Product{&domain.Product{
			GTIN:        "00031200452009",
			Description: "Ocean Spray Cranberry Classic Juice Drink 1 Litre",
			Contents: &struct {
				Quantity    float64 `json:"quantity,omitempty"`
				QuantityUOM string  `json:"quantityUom,omitempty"`
				AvgMeasure  string  `json:"avgMeasure,omitempty"`
				NetContents string  `json:"netContents,omitempty"`
			}{
				Quantity:    float64(1.0),
				QuantityUOM: "l",
				AvgMeasure:  "Average Measure (e)",
				NetContents: "1l ℮",
			},
			ProductCharacteristics: &struct {
				IsFood  bool `json:"isFood"`
				IsDrink bool `json:"isDrink"`
			}{
				IsFood:  false,
				IsDrink: true,
			},
		}}
		expected := &pb.ProductsResponse{
			Products: []*pb.Product{
				&pb.Product{
					Gtin:        "00031200452009",
					Description: "Ocean Spray Cranberry Classic Juice Drink 1 Litre",
					Contents: &pb.ProductContents{
						Quantity:    float32(1.0),
						QuantityUom: "l",
						AvgMeasure:  "Average Measure (e)",
						NetContents: "1l ℮",
					},
					Characterisitics: &pb.ProductCharacterisitics{
						IsFood:  false,
						IsDrink: true,
					},
				},
			},
		}
		mockGetter.On("Get", "").Return(mockProducts, nil)
		handler := NewGetProductsHandler(mockGetter)
		rq := &pb.GetProductsRequest{Barcode: ""}

		By("making call")
		actual, err := handler.GetProducts(nil, rq)

		By("assert")
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})

	Context("when an error occurs on retrieving products", func() {
		It("should return the correct error", func() {
			By("setup")
			mockGetter := &mocks.ProductGetter{}
			expectedError := errors.New("error retrieving products")
			mockGetter.On("Get", "").Return(nil, errors.New("mock error"))
			handler := NewGetProductsHandler(mockGetter)
			rq := &pb.GetProductsRequest{Barcode: ""}

			By("making call")
			actual, err := handler.GetProducts(nil, rq)

			By("assert")
			Expect(actual).To(BeNil())
			Expect(err).To(Equal(expectedError))
		})
	})
})
