package tests

import(
  "net/http"
	"net/http/httptest"
	. "licorne"
	"licorne/utilities"
	"io"
)

func Request(method string, route string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, route, nil)
	n := BuildMiddlewares()
	n.ServeHTTP(w, r)
	return w

}

// Set the database to licorne-test
func SetTestDatabase() {
	config := utilities.GetConfig()
	config.Database = "licorne-test"
}

func RequestWithBody(method string, route string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, route, body)
	BuildRouter().ServeHTTP(w, r)
	return w
}
