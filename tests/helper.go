package tests

import(
  "net/http"
	"net/http/httptest"
	. "licorne"
	"github.com/gin-gonic/gin"
	"licorne/utilities"
	"io"
)

var engine *gin.Engine = nil

func LicorneEngine() *gin.Engine {
	if engine == nil {
		engine = BuildEngine()
	}
	return engine
}

func Request(method string, route string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, route, nil)
	LicorneEngine().ServeHTTP(w, r)
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
	LicorneEngine().ServeHTTP(w, r)
	return w
}
