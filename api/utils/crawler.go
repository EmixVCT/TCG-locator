package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"tcglocator/model"
	"time"

	"github.com/robfig/cron"
)

func InitCrawl() {
	config := GetConfig()

	CrawlAll()

	c := cron.New()

	if config.CronCrawler != "" {
		err := c.AddFunc(config.CronCrawler, CrawlAll)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Crawler started (cron ", config.CronCrawler, ")")
		c.Start()
	}
}

func CrawlAll() {
	collections := GetCollections()

	for _, c := range collections {
		Debug("Crawl collection:", c.Name)
		for iName, i := range c.Items {
			Debug("Crawl items:", iName)
			for _, l := range i {

				//start crawling
				errCrawling := false
				if err := Crawl(l.Stock); err != nil {
					errCrawling = true
				}
				if err := Crawl(l.Price); err != nil {
					errCrawling = true
				}

				if !errCrawling {
					l.LastFetchDate = time.Now()
					l.State = true
					if l.Stock.Value != nil && l.Stock.Value.(bool) {
						l.LastStockDate = time.Now()
					}
				} else {
					l.State = false
				}

				Debug(l.Url)
				Debug("[CRAWLER] Price:", l.Price.Value, l.Currency, " Stock:", l.Stock.Value)
			}
			Debug("")
		}
	}
}

func Crawl(crawler *model.Crawler) error {

	if !crawler.IsApi {
		req, err := http.NewRequest("GET", crawler.Url, nil)
		if err != nil {
			return err
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		r := regexp.MustCompile(crawler.Regex)
		crawledValue := r.FindStringSubmatch(string(body))

		//Debug(len(crawledValue))
		if len(crawledValue) == 2 {
			//Debug(crawledValue[1])
			if crawler.Type == "string" {
				crawler.Value = crawledValue[1]
			} else if crawler.Type == "bool" {
				crawler.Value = true
			}
		} else {
			if crawler.Type == "bool" {
				crawler.Value = false
			} else if crawler.Type == "string" {
				crawler.Value = "-"
			}
		}
	} else {
		Debug("Not Implemented")
	}

	return nil
}
