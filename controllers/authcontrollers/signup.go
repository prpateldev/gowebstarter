package authcontrollers

import (
	"net/http"

	"gowebstarter/models/usermodel"
	"gowebstarter/viewmodels/authviewmodels"
	"encoding/json"
	"log"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &authviewmodels.Credentials{}

	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	// Next, insert the username, along with the hashed password into the database

	user := usermodel.User{
		uuid.New(),
		creds.Username,
		string(hashedPassword),
	}

	res, err := usermodel.CreateUser(user)
	if err != nil || res == false {
		// If there is any issue with inserting into the database, return a 500 error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	savedUser, err := usermodel.GetByUsername(user.Username)
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userJson, err := json.Marshal(savedUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
	// w.Write([]byte(fmt.Sprintf("signup")))
}
