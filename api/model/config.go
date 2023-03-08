package model

type Config struct {
	Version     string
	Debug       bool
	CronCrawler string
	SitesConfig []string

	Sites    []*Site
	Locators map[string]*Locator //url name Locator{}
}
