package controller

import (
	"log"
	"net/http"
	"tcglocator/utils"

	"github.com/go-chi/render"
)

func Locator(w http.ResponseWriter, r *http.Request) {
	locators := utils.GetConfig().Locators
	log.Println("Number of Items:", len(locators))
	render.Render(w, r, utils.Success("", utils.LocatorsRender(locators)))
}
