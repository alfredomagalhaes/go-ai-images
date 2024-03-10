package handler

import (
	"net/http"

	"github.com/alfredomagalhaes/go-ai-images/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {

	return home.Index().Render(r.Context(), w)
}
