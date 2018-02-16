package omnigo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOmnigo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Omnigo Suite")
}
