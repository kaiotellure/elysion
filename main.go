package main

import (
	"net/http"

	"github.com/kaiotellure/lysion/handlers"
	"github.com/kaiotellure/lysion/helpers"
	"github.com/kaiotellure/lysion/services/database"
)

func main() {
	database.Setup(helpers.Env(helpers.DATABASE))
	handlers.Setup(helpers.Env(helpers.PUBLIC_FOLDER))

	address := ":" + helpers.Env(helpers.PORT)
	if err := http.ListenAndServe(address, handlers.Router); err != nil {
		panic(err)
	}
}
