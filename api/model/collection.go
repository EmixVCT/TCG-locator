package model

import "net/http"

type Collection struct {
	Name     string
	Locators map[string]*Locator
}

func (rd *Collection) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	return nil
}
