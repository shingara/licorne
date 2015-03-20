package controllers_test

import (

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"licorne/models"
	"licorne/services"
	"licorne/tests"
	"licorne/utilities"
	"testing"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var GameFactory models.Game

var _ = BeforeEach(func() {
	tests.SetTestDatabase()

	_, GameFactory := tests.CreateFactoryGame()
	_, errs := services.CreateGame(&GameFactory)
	if (len(errs) > 0) {
		panic(errs)
	}
})

var _ = AfterEach(func() {
	utilities.DropDatabase()
})
