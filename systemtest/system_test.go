// +build system

package system_test

import (
	"context"

	pb "github.com/eggsbenjamin/groceries-tesco/grpc"
	"google.golang.org/grpc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("System", func() {
	It("should return the correct payload", func() {
		By("setup")
		barcode := "5012583200789"
		expected := &pb.ProductsResponse{
			Products: []*pb.Product{
				&pb.Product{
					Gtin:        "05012583200789",
					Description: "Right Guard Total Defence 5 Fresh Antiperspirant Deodorant 250Ml",
					Contents: &pb.ProductContents{
						Quantity:    float32(250.0),
						QuantityUom: "ml",
						AvgMeasure:  "Average Measure (e)",
						NetContents: "250ml ℮",
					},
					Characterisitics: &pb.ProductCharacterisitics{
						IsFood:  false,
						IsDrink: false,
					},
				},
			},
		}
		conn, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
		defer conn.Close()
		c := pb.NewTescoServiceClient(conn)

		By("making call")
		actual, err := c.GetProducts(context.Background(), &pb.GetProductsRequest{Barcode: barcode})

		By("assert")
		Expect(err).NotTo(HaveOccurred())
		Expect(actual).To(Equal(expected))
	})
})
