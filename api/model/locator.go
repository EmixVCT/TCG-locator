package model

import "net/http"

type Locator struct {
	Url           string
	Label         string
	Country       string
	State         bool
	LastFetchDate string
	LastStockDate string
	Stock         bool
	Price         string
	Currency      string
}

func (rd *Locator) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
