package authcontrollers

import (
	"gowebstarter/configs/cacheconfig"
	"gowebstarter/utils/cacheutils"
	"fmt"
	"net/http"
	"time"
)

func RefreshSessionTokenController(w http.ResponseWriter, r *http.Request) {
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

	// Now, create a new session token for the current user
	newSessionToken, err := cacheutils.SetSession(fmt.Sprintf("%s", session))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Delete the older session token
	_, err = cacheutils.DeleteSession(sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the new token as the users `session_token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: time.Now().Add(time.Duration(cacheconfig.CacheConfig.SESSION_TIME_OUT) * time.Second),
	})

}
