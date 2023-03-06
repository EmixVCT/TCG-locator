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

	if collection == "" || collections[collection] == nil || len(collections[collection].Items) == 0 {
		render.Render(w, r, utils.Err(errors.New("collection not found"), http.StatusBadRequest))
		return
	}

	log.Println("Loading collection:", collection)
	log.Println("Number of Items:", len(collections[collection].Items))
	s := 0
	for _, i := range collections[collection].Items {
		s += len(i)
	}
	log.Println("Number of Site:", s)

	render.Render(w, r, utils.Success("", utils.CollectionRender(collections[collection])))
}
