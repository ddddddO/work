package main

import (
	"github.com/ddddddO/work/go/gof/template"
)

func main() {
	abstractX := template.NewAbstract(
		template.NewSiteXCrawler("site-xxx", "https://xxxx.com"),
	)
	abstractY := template.NewAbstract(
		template.NewSiteXCrawler("site-yyy", "https://yyyy.com"),
	)

	if err := abstractX.Execute(); err != nil {
		panic(err)
	}
	if err := abstractY.Execute(); err != nil {
		panic(err)
	}

	// Output:
	// 2021/04/10 17:52:15 start site-xxx crawl.
	// Get request: https://xxxx.com
	// Scraping now...
	// Stored!
	// 2021/04/10 17:52:15 end site-xxx crawl.
	//
	// 2021/04/10 17:52:15 start site-yyy crawl.
	// Get request: https://yyyy.com
	// Scraping now...
	// Stored!
	// 2021/04/10 17:52:15 end site-yyy crawl.
	//
}