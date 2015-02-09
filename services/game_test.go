package services_test


import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
		. "licorne/services"
	"gopkg.in/mgo.v2/bson"
	)


var _ = Describe("AllGame", func(){

		It("return list of Game", func(){
			games, err := AllGame()
			Expect(err).To(BeNil())
			Expect(len(games)).To(Equal(1))
			Expect(games[0].Name).To(Equal("hello"))
			Expect(games[0].ID).To(Equal(Game_attr["id"].(bson.ObjectId)))
		})

})

var _ = Describe("GetGame", func(){

		It("return the game we request", func(){
			game, err := GetGame(Game_attr["id"])
			Expect(err).To(BeNil())
			Expect(game.Name).To(Equal("hello"))
			Expect(game.ID).To(Equal(Game_attr["id"].(bson.ObjectId)))
		})

})
