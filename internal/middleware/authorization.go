package middleware

import (
	"errors"
	"net/http"

	"github.com/AbeOwlu/ring-api/api"
	"github.com/AbeOwlu/ring-api/internal/tools"
	"go.uber.org/zap"
)

var UnAuthorizedError error = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggInit := zap.NewExample()
		logger := loggInit.Sugar()
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		// content := r.ParseForm()
		// chk := r.PostFormValue("api")

		if username == "" || token == "" {
			logger.Info(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w, err)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			logger.Info(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
		}

		next.ServeHTTP(w, r)
	})
}
