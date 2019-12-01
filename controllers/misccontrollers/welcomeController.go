package misccontrollers

import (
	"gowebstarter/utils/cacheutils"
	"fmt"
	"net/http"
)

func WelcomeController(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("session_token")

	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// For any other type of error return bad request
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken := c.Value
	// We then get the name of the user from our cache, where we set the session token
	session, err := cacheutils.GetSession(sessionToken)

	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if session == nil {
		// If the session token is not present in cache, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Finally, return the welcome message to the user
	w.Write([]byte(fmt.Sprintf("Welcome %s!", session)))

}
