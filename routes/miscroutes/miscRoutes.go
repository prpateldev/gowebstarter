package miscroutes

import (
	"gowebstarter/controllers/misccontrollers"
	"gowebstarter/controllers/pingcontroller"
	"gowebstarter/utils/routerutils"
)

func Ping() {
	routerutils.GetRouter().HandleFunc("/ping", pingcontroller.Ping)
}

func Welcome() {
	routerutils.GetRouter().HandleFunc("/welcome", misccontrollers.WelcomeController)
}
