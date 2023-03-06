package model

import (
	"net/http"
	"time"
)

type Locator struct {
	Url      string
	Country  string
	Currency string

	State         bool
	LastFetchDate time.Time
	LastStockDate time.Time
	Stock         *Crawler
	Price         *Crawler
}

func (rd *Locator) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
