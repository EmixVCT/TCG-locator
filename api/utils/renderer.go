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
	for iName, items := range collection.Items {
		for url, locator := range items {
			item := &model.ItemResponse{
				Url:           url,
				Label:         iName,
				Country:       locator.Country,
				Language:      locator.Language,
				Currency:      locator.Currency,
				State:         locator.State,
				LastFetchDate: locator.LastFetchDate,
				LastStockDate: locator.LastStockDate,
				Stock:         locator.Stock.Value.(bool),
				Price:         locator.Price.Value.(string),
			}
			list = append(list, item)
		}
	}
	return list
}
