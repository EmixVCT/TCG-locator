package controller

import (
	"errors"
	"net/http"
	"tcglocator/utils"

	"github.com/go-chi/render"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, utils.Err(errors.New("route does not exist"), http.StatusNotFound))

}
func NotValid(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, utils.Err(errors.New("method is not valid"), http.StatusMethodNotAllowed))
}
