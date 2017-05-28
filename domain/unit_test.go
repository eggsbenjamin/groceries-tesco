// +build unit

package domain_test

import (
	"errors"

	. "github.com/eggsbenjamin/groceries-tesco/domain"
	"github.com/eggsbenjamin/groceries-tesco/domain/mocks"
	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Domain Unit Tests", func() {
	var _ = Describe("ProductHandler", func() {
		It("returns the products correctly", func() {
			By("setup")
			mockClient := &mocks.TescoClient{}
			expected := []*Product{}
			mockClient.On("GetProductsByGTIN", mock.AnythingOfType("string")).Return(expected, nil)
			pH := NewProductHandler(mockClient)

			By("making call")
			actual, err := pH.Get("")

			By("assert")
			Expect(err).To(BeNil())
			Expect(actual).To(Equal(expected))
		})

		Context("on receiving an error from the tesco client", func() {
			It("it returns the correct error", func() {
				By("setup")
				mockClient := &mocks.TescoClient{}
				expectedError := errors.New("error")
				mockClient.On("GetProductsByGTIN", mock.AnythingOfType("string")).Return(nil, expectedError)
				pH := NewProductHandler(mockClient)

				By("making call")
				actual, err := pH.Get("")

				By("assert")
				Expect(actual).To(BeNil())
				Expect(err).To(Equal(expectedError))
			})
		})
	})
})
