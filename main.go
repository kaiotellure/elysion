package main

import (
	"net/http"

	"github.com/ikaio/tailmplx/internal/database"
	"github.com/ikaio/tailmplx/internal/help"
	"github.com/ikaio/tailmplx/internal/router"
)

func main() {
	database.Setup(help.Env(help.DATABASE, "tmp/main.development.db"))
	router.Setup(help.Env(help.PUBLIC_FOLDER, "web/public"))
	router.SetupRoutes()

	address := ":" + help.Env(help.PORT, "3000")
	if err := http.ListenAndServe(address, router.Router); err != nil {
		panic(err)
	}
}
