package main

import "fmt"
import "gowebstarter/configs/serverconfig"
import "gowebstarter/configs/dbconfig"
import "gowebstarter/configs/cacheconfig"
import "gowebstarter/utils/dbutils"
import "gowebstarter/utils/cacheutils"
import "gowebstarter/utils/routerutils"
import "gowebstarter/routes"
import "net/http"
import "time"

import "log"

func main() {

	// connecting to database
	dbutils.Init(dbconfig.DBConfig)
	// connectng to cache server
	cacheutils.Init(cacheconfig.CacheConfig)

	// creating the router instance
	routerutils.Init()

	// initiating route handlers
	routes.Init()

	serConfig := serverconfig.GetConfig()
	fmt.Println("starting crappbook api server on port", serConfig.PORT)
	server := http.Server{
		Handler: routerutils.GetRouter(),
		Addr:    "localhost:" + serConfig.PORT,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: time.Duration(serConfig.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(serConfig.ReadTimeout) * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// http.HandleFunc("/signin", handlers.Signin)
// 	http.HandleFunc("/signup", handlers.Signup)
// 	http.HandleFunc("/welcome", handlers.Welcome)
// 	http.HandleFunc("/refresh", handlers.RefreshSessionToken)
