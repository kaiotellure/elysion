package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kaiotellure/lysion/components"
	"github.com/kaiotellure/lysion/services/table"
)

var MEALS []components.Meal = []components.Meal{
	components.Meal{
		ID:   "pizza-pepperoni",
		Name: "Pizza de Pepperoni", Image: "https://receitasdepizza.com.br/wp-content/uploads/2023/02/Pizza-pizza-americana-com-pepperoni.webp",
		Description:  "A classica pizza italiana com o delicioso queijo mussarela recheada com o salame romano.",
		AllergyWarns: "Amendoin, Lactose",
		Price:        6900,
	},
}

func ListMeals() []components.Meal {
	return MEALS
}

func FindMealByID(id string) *components.Meal {
	for _, meal := range MEALS {
		if meal.ID == id {
			return &meal
		}
	}
	return nil
}

func handleMeal(w http.ResponseWriter, r *http.Request) {
	credential := getCredential(r)
	meal := FindMealByID(chi.URLParam(r, "id"))
	if meal == nil {
		Router.NotFoundHandler().ServeHTTP(w, r)
		return
	}

	components.Document(
		components.PageProps{
			Request: r, Auth: credential,
			Title: meal.Name,
		},
		components.PageMeal(*meal, table.ContainsItem(credential.Sub, meal.ID)),
	).Render(r.Context(), w)
}
