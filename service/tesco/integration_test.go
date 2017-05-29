// +build integration

package tesco_test

import (
	"net/http"

	"github.com/eggsbenjamin/groceries-tesco/domain"
	. "github.com/eggsbenjamin/groceries-tesco/service/tesco"
	"github.com/spf13/viper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tesco")
}

var _ = Describe("Tesco Service", func() {
	Context("when calling the Tesco API", func() {
		It("receives the expected products", func() {
			By("setup")
			client := &http.Client{}
			tescoSrv := NewTescoService(client)
			expected := []*domain.Product{}

			By("making call")
			actual, err := tescoSrv.GetProductsByGTIN("fake")

			By("assert")
			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})
	})
})
