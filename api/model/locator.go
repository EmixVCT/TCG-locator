package model

import (
	"net/http"
	"time"
)

type Locator struct {
	Label    string
	Country  string
	Currency string
	Language string

	State         bool
	LastFetchDate time.Time
	LastStockDate time.Time
	Stock         bool
	Price         string
}

func (rd *Locator) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
