package main

import (
	"net/http"

	"github.com/ikaio/tailmplx/database"
	"github.com/ikaio/tailmplx/help"
	"github.com/ikaio/tailmplx/router"
)

func main() {
	database.Setup(help.Env(help.DATABASE, "tmp/main.test.db"))

	router.Setup(help.Env(help.PUBLIC_FOLDER, "web/public"))
	router.SetupRoutes()

	address := ":"+help.Env(help.PORT, "3000")
	if err := http.ListenAndServe(address, router.Router); err != nil {
		panic(err)
	}
}
