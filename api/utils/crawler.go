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
	config := GetConfig()
	for _, s := range config.Sites {

		Debug("\n", "Crawl site:", s.Name, "\n")

		for name, l := range s.Urls {
			config.Locators[l.Url] = &model.Locator{}

			config.Locators[l.Url].Label = name
			config.Locators[l.Url].Country = s.Country
			config.Locators[l.Url].Vendor = s.Name
			config.Locators[l.Url].Currency = s.Currency
			config.Locators[l.Url].Language = l.Language

			if v, err := Crawl(s.Price, l.Url); err == nil {
				config.Locators[l.Url].Price = v.(string)
				if v, err := Crawl(s.Stock, l.Url); err == nil {
					config.Locators[l.Url].Stock = v.(bool)

					config.Locators[l.Url].LastFetchDate = time.Now()
					config.Locators[l.Url].State = true
					if config.Locators[l.Url].Stock {
						config.Locators[l.Url].LastStockDate = time.Now()
					} else {
						config.Locators[l.Url].Stock = false
					}
					Debug(l.Url)
					Debug("[CRAWLER] Price:", config.Locators[l.Url].Price, config.Locators[l.Url].Currency, " Stock:", config.Locators[l.Url].Stock, "\n")
					continue
				}
			}
			Debug(l.Url, "not continue")
			config.Locators[l.Url].State = false
		}
	}

}

func Crawl(crawler *model.Crawler, url string) (interface{}, error) {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)

	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(crawler.Regex)
	crawledValue := r.FindStringSubmatch(string(body))

	if len(crawledValue) == 2 {
		if crawler.Type == "bool" {
			return true, nil
		}
		if crawledValue[1] == "" {
			return "-", nil
		}

		return crawledValue[1], nil

	}
	if crawler.Type == "bool" {
		return false, nil
	}
	return "-", nil

}
