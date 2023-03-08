package model

type LocatorInfo struct {
	Language string
	Url      string
}

type Site struct {
	Name     string
	Country  string
	Currency string
	Stock    *Crawler
	Price    *Crawler
	Urls     map[string]*LocatorInfo
}
