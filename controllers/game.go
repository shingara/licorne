package controllers

import (
	"net/http"
	"fmt"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home")
}
