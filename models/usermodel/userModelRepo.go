package usermodel

import "gowebstarter/utils/dbutils"

func GetPasswordByUsername(username string) (string, error) {
	// get db instance
	db := dbutils.GetDB()
	user := &User{}
	// query to find user by username
	err := db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&user.Password)

	return user.Password, err
}

func GetByUsername(username string) (*User, error) {
	// get db instance
	db := dbutils.GetDB()
	user := &User{}
	// query to find user by username
	err := db.QueryRow("SELECT `id`, `username`, `password` FROM users WHERE username=?", username).Scan(&user.Id, &user.Username, &user.Password)

	return user, err
}

func CreateUser(user User) (bool, error) {
	// get db instance
	db := dbutils.GetDB()
	// create statement to insert user
	stmt, err := db.Prepare("INSERT INTO users (id, username, password, created_at) values(?, ?, ?, NOW())")
	if err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		return false, err
	}
	_, err = stmt.Exec(user.Id, user.Username, user.Password)
	if err != nil {
		return false, err
	}
	return true, err
}
