package controllers_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
	. "licorne/tests"
)

var _ = Describe("GameController", func() {

	// var (
	// 	body []byte
	// 	err error
	// 	//todos []Todo
	// )

	Context("List all games", func() {
		It("returns a 200 Status Code", func() {
			response := Request("GET", "/games")
			Expect(response.Code).To(Equal(200))
		})
	})

})
