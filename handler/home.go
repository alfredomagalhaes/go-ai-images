package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/alfredomagalhaes/go-ai-images/db"
	"github.com/alfredomagalhaes/go-ai-images/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	account, err := db.GetAccountByUserID(user.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	fmt.Printf("%+v\n", account)
	return home.Index().Render(r.Context(), w)
}
