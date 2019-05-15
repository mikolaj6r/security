package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeResult struct {
	mail string
}

func (item ScrapeResult) toString() string {
	return fmt.Sprintf("%s", item.mail)
}

type Parser interface {
	ParsePage(*goquery.Document) []ScrapeResult
}

func getRequest(url string) (*http.Response, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

type ArrayWithId struct {
	array []ScrapeResult
	id    int
}

func crawlPage(targetURL string, num int, parser Parser) ([]ScrapeResult, int) {

	fmt.Println("Requesting: ", targetURL)
	resp, _ := getRequest(targetURL)

	doc, _ := goquery.NewDocumentFromResponse(resp)
	pageResults := parser.ParsePage(doc)

	return pageResults, num
}

type boolChan chan bool

func Crawl(startURL string, parser Parser, first int, num int) []ArrayWithId {
	var wg sync.WaitGroup
	arr := make([]ArrayWithId, num+1)

	wg.Add(num)
	for i := first; i < (first + num); i++ {
		var link = startURL + strconv.Itoa(i)
		go func(link string, index int, parser Parser) {
			pageResults, num := crawlPage(link, index, parser)
			arr[index] = ArrayWithId{pageResults, num}
			wg.Done()
		}(link, i, parser)
	}
	wg.Wait()
	return arr
}
