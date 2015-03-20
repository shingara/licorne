package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"licorne/models"
	"licorne/tests"
	"licorne/utilities"
	"licorne/services"

	"testing"
)

func TestServices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite")
}

var (
	Game_attr map[string]interface{}
	GameFactory models.Game
)

var _ = BeforeSuite(func() {
	tests.SetTestDatabase()
})

var _ = BeforeEach(func(){
	Game_attr, GameFactory  = tests.CreateFactoryGame()
	services.CreateGame(&GameFactory)
})

var _ = AfterEach(func() {
	utilities.DropDatabase()
})
