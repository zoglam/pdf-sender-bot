package rest

import (
	"net/http"
)

func (a *Rest) PostProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(http.StatusOK)
}
