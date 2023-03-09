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

func LocatorsRender(locators map[string]*model.Locator) []render.Renderer {
	list := []render.Renderer{}
	for url, locator := range locators {
		item := &model.ItemResponse{
			Url:           url,
			Label:         locator.Label,
			Country:       locator.Country,
			Vendor:        locator.Vendor,
			Language:      locator.Language,
			Currency:      locator.Currency,
			State:         locator.State,
			LastFetchDate: locator.LastFetchDate,
			LastStockDate: locator.LastStockDate,
			Stock:         locator.Stock,
			Price:         locator.Price,
		}
		list = append(list, item)
	}
	return list
}
