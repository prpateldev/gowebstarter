package authcontrollers

import (
	"gowebstarter/configs/cacheconfig"
	"gowebstarter/models/usermodel"
	"gowebstarter/utils/cacheutils"
	"gowebstarter/viewmodels/authviewmodels"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	creds := &authviewmodels.Credentials{}

	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := usermodel.GetByUsername(creds.Username)

	if err != nil {
		log.Println(err)
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent

	// Set Session
	sessionToken, err := cacheutils.SetSession(creds.Username)
	if err != nil {
		// If there is an error in setting the cache, return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds, the same as the cache
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: time.Now().Add(time.Duration(cacheconfig.CacheConfig.SESSION_TIME_OUT) * time.Second),
	})

}
