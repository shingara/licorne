package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"licorne/tests"

	"encoding/json"
)

var _ = Describe("GameController", func() {

	Context("List all games", func() {
		It("returns a 200 Status Code", func() {
			response := tests.Request("GET", "/v1/games")
			Expect(response.Code).To(Equal(200))
			content_type := response.HeaderMap.Get("Content-Type")
			Expect(content_type).To(Equal("application/json; charset=utf-8"))


			var json_return map[string]interface{}
			err := json.Unmarshal(response.Body.Bytes(), &json_return)
			Expect(err).To(BeNil())
			Expect(json_return["data"]).NotTo(BeNil())
			games := json_return["data"].([]interface{})
			Expect(len(games)).To(Equal(1))
		})
	})

	Context("Get a games", func() {

		It("returns a 200 Status Code", func(){
			id := GameFactory.ID.Hex()
			url := "/v1/games/" + id
			response := tests.Request("GET", url)
			Expect(response.Code).To(Equal(200))
			content_type := response.HeaderMap.Get("Content-Type")
			Expect(content_type).To(Equal("application/json; charset=utf-8"))

			var json_return map[string]interface{}
			err := json.Unmarshal(response.Body.Bytes(), &json_return)
			Expect(err).To(BeNil())
			Expect(json_return["type"]).To(Equal("games"))
			Expect(json_return["id"]).To(Equal(id))

		})
	})

})
