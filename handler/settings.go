package handler

import (
	"net/http"

	"github.com/alfredomagalhaes/go-ai-images/view/settings"
)

func HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	return render(w, r, settings.Index(user))
}
