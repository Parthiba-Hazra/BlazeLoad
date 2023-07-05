package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func crawlWebsite(urlString string, visited map[string]bool) {

	parsedURL, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Get(parsedURL.String())
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parsedURL.String())

	doc.Find("a").Each(func(index int, element *goquery.Selection) {

		href, exists := element.Attr("href")
		if exists {

			absoluteURL := parsedURL.ResolveReference(&url.URL{Path: href}).String()

			if !visited[absoluteURL] {
				visited[absoluteURL] = true
				crawlWebsite(absoluteURL, visited)
			}
		}
	})
}

func main() {
	startURL := "https://jsonplaceholder.typicode.com"

	visited := make(map[string]bool)

	crawlWebsite(startURL, visited)
}
