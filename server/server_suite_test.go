// +build unit

package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTesco(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Unit Test Suite")
}
