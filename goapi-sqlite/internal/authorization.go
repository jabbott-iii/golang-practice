package internal

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError error = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *DatabaseInterface
		database, err = NewDatabase()
		if err != nil {
			InternalErrorHandler(w)
			return
		}

		var loginDetails *LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
