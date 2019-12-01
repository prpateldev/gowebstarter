package authroutes

import (
	"gowebstarter/controllers/authcontrollers"
	"gowebstarter/utils/routerutils"
)

func Signup() {
	routerutils.GetRouter().HandleFunc("/signup", authcontrollers.Signup)
}

func Signin() {
	routerutils.GetRouter().HandleFunc("/signin", authcontrollers.Signin)
}

func RefreshAuth() {
	routerutils.GetRouter().HandleFunc("/auth/refresh", authcontrollers.Signin)
}
