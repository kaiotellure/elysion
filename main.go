package main

import (
	"net/http"

	"github.com/ikaio/tailmplx/handlers"
	"github.com/ikaio/tailmplx/helpers"
	"github.com/ikaio/tailmplx/services/database"
)

func main() {
	database.Setup(helpers.Env(helpers.DATABASE))
	handlers.Setup(helpers.Env(helpers.PUBLIC_FOLDER))

	address := ":" + helpers.Env(helpers.PORT)
	if err := http.ListenAndServe(address, handlers.Router); err != nil {
		panic(err)
	}
}
