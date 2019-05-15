package main

import (
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DummyParser struct {
}

func (d DummyParser) ParsePage(doc *goquery.Document) []ScrapeResult {
	var results []ScrapeResult

	re := regexp.MustCompile(`\/people\/details\/(.*)`)
	items := doc.Find(".content .table .table__link")
	items.Each(func(i int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		mail := strings.TrimSpace(re.ReplaceAllString(href, "$1@put.poznan.pl"))

		results = append(results, ScrapeResult{mail})
	})
	return results
}

func main() {
	start := 1
	end := 18
	f, _ := os.Create("mails.txt")
	defer f.Close()
	d := DummyParser{}
	res := Crawl("https://sin.put.poznan.pl/people/all?page=", d, start, end)

	for i := start; i < (start + end); i++ {
		for _, el := range res[i].array {
			f.WriteString(el.mail + ";")
		}
		f.WriteString("\n\n\n")
	}
	f.Sync()
}
