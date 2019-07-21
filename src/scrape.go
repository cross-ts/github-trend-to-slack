package main

import (
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

func makeTrendsSelectionFrom(response *http.Response) *goquery.Selection {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	trendSelections := doc.Find(TREND_QUERY)
	return trendSelections
}

func makeGithubTrendFrom(s *goquery.Selection) (trend GithubTrend) {
	repo, exists := s.Find("h1 > a").Attr("href")
	if !exists {
		panic("Not found!")
	}
	desc := s.Find("p").Text()
	trend = GithubTrend{repo: repo, desc: desc}
	return
}

func makeGithubTrendsFrom(s *goquery.Selection) (result []GithubTrend) {
	s.Each(func(i int, s *goquery.Selection) {
		trend := makeGithubTrendFrom(s)
		result = append(result, trend)
	})
	return
}

func Scrape() (trends []GithubTrend) {
	res := httpRequest()
	defer res.Body.Close()
	selections := makeTrendsSelectionFrom(res)
	trends = makeGithubTrendsFrom(selections)
	return
}
