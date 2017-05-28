// +build unit

package tesco_test

import (
	"github.com/eggsbenjamin/groceries-tesco/domain"
	. "github.com/eggsbenjamin/groceries-tesco/service/tesco"
	"github.com/eggsbenjamin/groceries-tesco/service/tesco/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Tesco Unit Tests", func() {
	var _ = Describe("Tesco Service", func() {
		It("returns products correctly", func() {
			By("setup")
			mockClient := &mocks.HTTPClient{}
			mockBody := `{"products": [{"gtin": "00031200452009","description": "Ocean Spray Cranberry Classic Juice Drink 1 Litre", "qtyContents": { "quantity": 1.0, "quantityUom": "l", "avgMeasure": "Average Measure (e)", "netContents": "1l ℮" }, "productCharacteristics": { "isFood": false, "isDrink": true }}]}`
			mockResponse := CreateMockResponse(200, mockBody)
			expected := []*domain.Product{&domain.Product{
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
			tescoSrv := NewTescoService(mockClient)
			mockClient.On("Do", mock.Anything).Return(mockResponse, nil)

			By("making call")
			actual, err := tescoSrv.GetProductsByGTIN("123")

			By("assert")
			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		Context("when no products are returned", func() {
			It("returns an empty array", func() {
				By("setup")
				mockClient := &mocks.HTTPClient{}
				mockResponse := CreateMockResponse(200, `{ "products": [] }`)
				expected := []*domain.Product{}
				tescoSrv := NewTescoService(mockClient)
				mockClient.On("Do", mock.Anything).Return(mockResponse, nil)

				By("making call")
				actual, err := tescoSrv.GetProductsByGTIN("123")

				By("assert")
				Expect(err).To(BeNil())
				Expect(actual).To(Equal(expected))
			})
		})
	})
})
