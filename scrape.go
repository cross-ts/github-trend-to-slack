package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const GITHUB_TRENDING_URL string = "https://github.com/trending"
const TREND_QUERY string = "html body div main div div.Box div article.Box-row"

func httpRequest() *http.Response {
	res, err := http.Get(GITHUB_TRENDING_URL)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		panic("github.com/trending didn't return code 200.")
	}
	return res
}

func pickTrendsFrom(response *http.Response) *goquery.Selection {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	trendSelections := doc.Find(TREND_QUERY)
	return trendSelections
}

func hoge(i int, s *goquery.Selection) {
	repo, _ := s.Find("h1 > a").Attr("href")
	desc := s.Find("p").Text()
	fmt.Println(repo)
	fmt.Println(desc)
}

func Scrape() {
	res := httpRequest()
	defer res.Body.Close()

	trends := pickTrendsFrom(res)
	trends.Each(hoge)
}
