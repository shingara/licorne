package client

import(
	"os"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func BaseUrl() string {
	host := os.Getenv("WEB_1_PORT_9000_TCP_ADDR")
	port := os.Getenv("WEB_1_PORT_9000_TCP_PORT")
	return "http://" + host + ":" + port
}

func GetContent(url string) (string, []error) {
	_, body, errs := gorequest.New().Get(url).End()
	// At this point we're done - simply return the bytes
	return body, errs
}

func PostContent(url string, params interface{}) (string, []error) {
	_, body, errs := gorequest.New().Post(url).
		Send(params).
		End()
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	return body, errs
}
