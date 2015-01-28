package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLicorne(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Licorne Suite")
}

