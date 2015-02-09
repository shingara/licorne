package controllers

import (
	"github.com/unrolled/render"
	"net/http"
)
var r = render.New()

func RenderJsonAPI(rw http.ResponseWriter, status int, v interface{}){
	r.JSON(rw, status, v)
}
