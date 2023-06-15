package controller

import (
	"net/http"
	"tcglocator/utils"

	"github.com/go-chi/render"
)

func Locator(w http.ResponseWriter, r *http.Request) {
	locators := utils.GetConfig().Locators
	render.Render(w, r, utils.Success("", utils.LocatorsRender(locators)))
}
