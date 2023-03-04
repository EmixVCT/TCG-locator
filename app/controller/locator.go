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
	config := utils.GetConfig()
	collection := chi.URLParam(r, "collection")

	if collection == "" || len(config.Collections[collection].Locators) == 0 {
		render.Render(w, r, utils.Err(errors.New("collection not found"), http.StatusBadRequest))
		return
	}

	log.Println("Loading collection:", collection)
	log.Println("Number of Locators:", len(config.Collections[collection].Locators))

	render.Render(w, r, utils.Success("", utils.CollectionRender(config.Collections[collection])))
}
