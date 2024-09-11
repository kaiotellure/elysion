package handlers

import (
	"net/http"

	"github.com/kaiotellure/lysion/components"
	"github.com/kaiotellure/lysion/services/table"
)

func handleTable(w http.ResponseWriter, r *http.Request) {
	credential := getCredential(r)
	if credential == nil {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}

	components.Document(
		components.PageProps{
			Request: r, Auth: credential,
			Title: "Sua Mesa",
		},
		components.PageTable(table.ListTable(credential.Sub)),
	).Render(r.Context(), w)
}

func handleTableAdd(w http.ResponseWriter, r *http.Request) {
	credential := getCredential(r)
	if credential == nil {
		components.GoogleError("not logged in").Render(r.Context(), w)
		return
	}

	id := r.URL.Query().Get("id")
	meal := FindMealByID(id)
	if meal == nil {
		components.GoogleError("could not find meal: "+id).Render(r.Context(), w)
		return
	}

	table.AddItem(credential.Sub, id)
	components.MealOrderButton("remove", id).Render(r.Context(), w)
}

func handleTableRemove(w http.ResponseWriter, r *http.Request) {
	credential := getCredential(r)
	if credential == nil {
		components.GoogleError("not logged in").Render(r.Context(), w)
		return
	}

	id := r.URL.Query().Get("id")
	meal := FindMealByID(id)
	if meal == nil {
		components.GoogleError("could not find meal: "+id).Render(r.Context(), w)
		return
	}

	table.RemoveItem(credential.Sub, id)
	components.MealOrderButton("add", id).Render(r.Context(), w)
}
