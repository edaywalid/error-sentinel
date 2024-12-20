package handlers

import (
	"net/http"
)

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("panic!")
}
