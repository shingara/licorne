package tests

import(
  "net/http"
	"net/http/httptest"
	. "licorne"
	"io"
)

func Request(method string, route string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, route, nil)

	BuildRouter().ServeHTTP(w, r)
	return w

}

func RequestWithBody(method string, route string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, route, body)
	BuildRouter().ServeHTTP(w, r)
	return w
}
