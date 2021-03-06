package controllers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"licorne/services"
	"licorne/models"
	"licorne/tests"
	"licorne/utilities"
	"testing"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controllers Suite")
}

var GameFactory models.Game

var _ = BeforeSuite(func() {
	tests.SetTestDatabase()

	_, GameFactory  = tests.CreateFactoryGame()
	services.CreateGame(&GameFactory)
})

var _ = AfterSuite(func() {
	utilities.DropDatabase()
})

