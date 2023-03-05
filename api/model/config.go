package model

type Config struct {
	Version     string
	Debug       bool
	CronCrawler string

	Collections map[string]*Collection
}
