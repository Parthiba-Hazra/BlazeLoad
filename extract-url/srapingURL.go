package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "http://google.com"

	var extractedURLs []string

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			if strings.HasPrefix(href, "http") {
				extractedURLs = append(extractedURLs, href)
			} else {
				absoluteURL := url + href
				extractedURLs = append(extractedURLs, absoluteURL)
			}
		}
	})

	doc.Find("link").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			if strings.HasPrefix(href, "http") {
				extractedURLs = append(extractedURLs, href)
			} else {
				absoluteURL := url + href
				extractedURLs = append(extractedURLs, absoluteURL)
			}
		}
	})

	doc.Find("script").Each(func(index int, element *goquery.Selection) {
		src, exists := element.Attr("src")
		if exists {
			if strings.HasPrefix(src, "http") {
				extractedURLs = append(extractedURLs, src)
			} else {
				absoluteURL := url + src
				extractedURLs = append(extractedURLs, absoluteURL)
			}
		}
	})

	for _, link := range extractedURLs {
		fmt.Println(link)
	}
}
