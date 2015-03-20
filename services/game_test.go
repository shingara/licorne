package services_test


import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "licorne/services"
	"licorne/models"

	"licorne/tests"

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
			game, err := GetGame(Game_attr["id"].(bson.ObjectId).Hex())
			Expect(err).To(BeNil())
			Expect(game.Name).To(Equal("hello"))
			Expect(game.ID).To(Equal(Game_attr["id"].(bson.ObjectId)))
		})

})

var _ = Describe("CreateGame", func(){

	It("create a game", func(){

		games, _ := AllGame()
		nb_game_before := len(games)

		_, game_factory := tests.CreateFactoryGame()
		_, errs := CreateGame(&game_factory)
		if (len(errs) > 0) {
			panic(errs)
		}

		games, _ = AllGame()
		nb_game_after := len(games)
		Expect(nb_game_after).To(Equal(nb_game_before + 1))

	})

	It("generate a ID if not define", func(){
		_, game_factory := tests.CreateFactoryGame()
		game, err:= CreateGame(&models.Game{
			Name: "hello",
			Players: game_factory.Players,
		})
		Expect(len(err)).To(Equal(0))
		Expect(game.ID).ToNot(BeNil())
		new_game, err_2 := GetGame(game.ID.Hex())
		Expect(err_2).To(BeNil())
		Expect(new_game).ToNot(BeNil())
	})

	It("can't save if no players attach less than 2", func(){
		_, err:= CreateGame(&models.Game{
			Name: "hello",
		})
		errs := []error{}
		errs = append(errs, errors.New("need at least 2 players"))
		Expect(err).To(Equal(errs))
	})

	It("can't save if no players attach more than 4", func(){
		playerList := make([]models.Player, 5)
		_, err:= CreateGame(&models.Game{
			Name: "hello",
			Players: playerList,
		})
		errs := []error{}
		errs = append(errs, errors.New("need less than 4 players"))
		Expect(err).To(Equal(errs))
	})

})
