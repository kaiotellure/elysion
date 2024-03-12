package main

import (
	"net/http"

	"github.com/ikaio/tailmplx/database"
	"github.com/ikaio/tailmplx/handlers"
	"github.com/ikaio/tailmplx/help"
)

func main() {
	database.Setup(help.Env(help.DATABASE, "tmp/main.development.db"))
	handlers.Setup(help.Env(help.PUBLIC_FOLDER, "web/public"))
	handlers.SetupRoutes()

	address := ":" + help.Env(help.PORT, "3000")
	if err := http.ListenAndServe(address, handlers.Router); err != nil {
		panic(err)
	}
}
