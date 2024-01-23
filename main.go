package main

import (
	"net/http"

	"github.com/ikaio/tailmplx/components"
	"github.com/ikaio/tailmplx/database"
	"github.com/ikaio/tailmplx/router"
	"github.com/ikaio/tailmplx/utilities"
)

func main() {
	database.Setup(utilities.Env("DATABASE", "tmp/main.test.db"))

	router.Setup(utilities.Env("PUBLIC_FOLDER", "web/public"))
	components.AttachHandlersFromPoolTo(router.Router)

	err := http.ListenAndServe(":"+utilities.Env("PORT", "3000"), router.Router)
	if err != nil {
		panic(err)
	}
}
