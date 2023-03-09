package model

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (s *SuccessResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, 200)
	s.Status = "success"

	return nil
}

type ItemResponse struct {
	Url      string
	Label    string
	Country  string
	Vendor   string
	Language string
	Currency string

	State         bool
	LastFetchDate time.Time
	LastStockDate time.Time
	Stock         bool
	Price         string
}

func (s *ItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, 200)
	return nil
}
