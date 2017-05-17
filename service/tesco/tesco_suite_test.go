package tesco_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTesco(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tesco Suite")
}
