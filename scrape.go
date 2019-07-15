package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func Scrape() {
	res, err := http.Get("https://github.com/trending")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	// Find the review items
	doc.Find("html > body > div > main > div > div.Box > div > article.Box-row").Each(func(i int, s *goquery.Selection) {
		repo, _ := s.Find("h1 > a").Attr("href")
		desc := s.Find("p").Text()
		fmt.Println(repo)
		fmt.Println(desc)
	})
}
