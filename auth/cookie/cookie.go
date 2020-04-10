// Package cookie provides standard CouchDB cookie auth as described at
// http://docs.couchdb.org/en/2.0.0/api/server/authn.html#cookie-authentication
package cookie

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-kivik/kivik/v4"
	"github.com/go-kivik/kivik/v4/errors"
	"github.com/go-kivik/kivikd/v4"
	"github.com/go-kivik/kivikd/v4/auth"
	"github.com/go-kivik/kivikd/v4/authdb"
	"github.com/go-kivik/kivikd/v4/cookies"
)

const typeJSON = "application/json"

// Auth provides CouchDB Cookie authentication.
type Auth struct{}

var _ auth.Handler = &Auth{}

// MethodName returns "cookie"
func (a *Auth) MethodName() string {
	return "cookie" // For compatibility with the name used by CouchDB
}

// Authenticate authenticates a request with cookie auth against the user store.
func (a *Auth) Authenticate(w http.ResponseWriter, r *http.Request) (*authdb.UserContext, error) {
	if r.URL.Path == "/_session" {
		switch r.Method {
		case http.MethodPost:
			return nil, postSession(w, r)
		case http.MethodDelete:
			return nil, deleteSession(w)
		}
	}
	return a.validateCookie(r)
}

func (a *Auth) validateCookie(r *http.Request) (*authdb.UserContext, error) {
	store := kivikd.GetService(r).UserStore
	cookie, err := r.Cookie(kivik.SessionCookieName)
	if err != nil {
		return nil, nil
	}
	name, _, err := cookies.DecodeCookie(cookie.Value)
	if err != nil {
		return nil, nil
	}
	user, err := store.UserCtx(r.Context(), name)
	if err != nil {
		// Failed to look up the user
		return nil, nil
	}
	s := kivikd.GetService(r)
	valid, err := s.ValidateCookie(user, cookie.Value)
	if err != nil || !valid {
		return nil, nil
	}
	return user, nil
}

func postSession(w http.ResponseWriter, r *http.Request) error {
	authData := struct {
		Name     *string `form:"name" json:"name"`
		Password string  `form:"password" json:"password"`
	}{}
	if err := kivikd.BindParams(r, &authData); err != nil {
		return errors.Status(http.StatusBadRequest, "unable to parse request data")
	}
	if authData.Name == nil {
		return errors.Status(http.StatusBadRequest, "request body must contain a username")
	}
	s := kivikd.GetService(r)
	user, err := s.UserStore.Validate(r.Context(), *authData.Name, authData.Password)
	if err != nil {
		return err
	}
	next, err := redirectURL(r)
	if err != nil {
		return err
	}

	// Success, so create a cookie
	token, err := s.CreateAuthToken(*authData.Name, user.Salt, time.Now().Unix())
	if err != nil {
		return err
	}
	w.Header().Set("Cache-Control", "must-revalidate")
	http.SetCookie(w, &http.Cookie{
		Name:     kivik.SessionCookieName,
		Value:    token,
		Path:     "/",
		MaxAge:   getSessionTimeout(s),
		HttpOnly: true,
	})
	w.Header().Add("Content-Type", typeJSON)
	if next != "" {
		w.Header().Add("Location", next)
		w.WriteHeader(http.StatusFound)
	}
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"ok":    true,
		"name":  user.Name,
		"roles": user.Roles,
	})
}

func redirectURL(r *http.Request) (string, error) {
	next, ok := kivikd.StringQueryParam(r, "next")
	if !ok {
		return "", nil
	}
	if !strings.HasPrefix(next, "/") {
		return "", errors.Status(http.StatusBadRequest, "redirection url must be relative to server root")
	}
	if strings.HasPrefix(next, "//") {
		// Possible schemaless url
		return "", errors.Status(http.StatusBadRequest, "invalid redirection url")
	}
	parsed, err := url.Parse(next)
	if err != nil {
		return "", errors.Status(http.StatusBadRequest, "invalid redirection url")
	}
	return parsed.String(), nil
}

func deleteSession(w http.ResponseWriter) error {
	http.SetCookie(w, &http.Cookie{
		Name:     kivik.SessionCookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	w.Header().Add("Content-Type", typeJSON)
	w.Header().Set("Cache-Control", "must-revalidate")
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"ok": true,
	})
}

func getSessionTimeout(s *kivikd.Service) int {
	if s.Conf().IsSet("couch_httpd_auth.timeout") {
		return s.Conf().GetInt("couch_httpd_auth.timeout")
	}
	return kivikd.DefaultSessionTimeout
}
