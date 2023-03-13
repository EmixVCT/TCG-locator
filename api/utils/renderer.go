package utils

import (
	"sort"
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

	keys := make([]string, 0, len(locators))
	for k, _ := range locators {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, url := range keys {
		item := &model.ItemResponse{
			Url:           url,
			Label:         locators[url].Label,
			Country:       locators[url].Country,
			Vendor:        locators[url].Vendor,
			Language:      locators[url].Language,
			Currency:      locators[url].Currency,
			State:         locators[url].State,
			LastFetchDate: locators[url].LastFetchDate,
			LastStockDate: locators[url].LastStockDate,
			Stock:         locators[url].Stock,
			Price:         locators[url].Price,
		}
		list = append(list, item)
	}

	return list
}
