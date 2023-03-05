package utils

import (
	"tcglocator/model"

	"github.com/go-chi/render"
)

func Err(err error, statusCode int) render.Renderer {
	return &model.ErrResponse{
		Err:            err,
		HTTPStatusCode: statusCode,
		Message:        err.Error(),
		Status:         "error",
	}
}

func Success(message string, data interface{}) render.Renderer {
	return &model.SuccessResponse{
		Message: message,
		Data:    data,
	}
}

func CollectionRender(collection *model.Collection) []render.Renderer {
	list := []render.Renderer{}
	for _, locator := range collection.Locators {
		list = append(list, locator)
	}
	return list
}
