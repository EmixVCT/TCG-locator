package controller

import (
	"errors"
	"log"
	"net/http"
	"tcglocator/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func Locator(w http.ResponseWriter, r *http.Request) {
	collections := utils.GetCollections()
	collection := chi.URLParam(r, "collection")

	if collection == "" || collections[collection] == nil || len(collections[collection].Locators) == 0 {
		render.Render(w, r, utils.Err(errors.New("collection not found"), http.StatusBadRequest))
		return
	}

	log.Println("Loading collection:", collection)
	log.Println("Number of Locators:", len(collections[collection].Locators))

	render.Render(w, r, utils.Success("", utils.CollectionRender(collections[collection])))
}
