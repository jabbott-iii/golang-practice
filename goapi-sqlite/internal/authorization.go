package internal

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		database, err := NewDatabase()
		if err != nil {
			log.Error(err)
			InternalErrorHandler(w)
			return
		}

		loginDetails := database.GetUserLoginDetails(username)
		if loginDetails != nil && token != loginDetails.AuthToken {
			log.Error(UnAuthorizedError)
			RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
