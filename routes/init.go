package routes

import "gowebstarter/routes/miscroutes"
import "gowebstarter/routes/authroutes"

func Init() {
	// list misc routes
	miscroutes.Ping()
	miscroutes.Welcome()

	// list auth routes
	authroutes.Signin()
	authroutes.Signup()
}
