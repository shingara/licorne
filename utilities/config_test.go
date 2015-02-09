package utilities_test


import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
		. "licorne/utilities"
		"os"
	)


var _ = Describe("Config", func(){

	Context("GetConfig()", func(){

		It("return defaut value", func(){
			config := GetConfig()
			Expect(config.Database).To(Equal("licorne-dev"))
			Expect(config.Host).To(Equal(os.Getenv("DB_1_PORT_27017_TCP_ADDR")))
			Expect(config.Port).To(Equal("27017"))
		})

	})
})
