package handler

import (
	"log/slog"
	"net/http"

	"github.com/alfredomagalhaes/go-ai-images/pkg/sb"
	"github.com/alfredomagalhaes/go-ai-images/pkg/util"
	"github.com/alfredomagalhaes/go-ai-images/view/auth"
	"github.com/nedpals/supabase-go"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, auth.Login())
}

func HandleLoginCreate(w http.ResponseWriter, r *http.Request) error {
	credentials := supabase.UserCredentials{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if !util.IsValidEmail(credentials.Email) {
		return render(w, r, auth.LoginForm(credentials, auth.LoginErrors{
			Email: "invalid email format, check and try again",
		}))
	}

	resp, err := sb.Client.Auth.SignIn(r.Context(), credentials)

	if err != nil {
		slog.Error("error while authenticating user", err)
		return render(w, r, auth.LoginForm(credentials, auth.LoginErrors{
			InvalidCredentials: "The provided credentials are invalid",
		}))
	}

	cookie := &http.Cookie{
		Value:    resp.AccessToken,
		Name:     "at",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "", http.StatusSeeOther)

	return nil
}
