package handler

import (
	"net/http"

	"github.com/alfredomagalhaes/go-ai-images/view/auth"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Login().Render(r.Context(), w)
}
