package main

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func main () {
	doc, err := goquery.NewDocument("https://twitter.com/okkun_sh")
	if err != nil {
		fmt.Println("error!!!")
	}

	selection := doc.Find(".js-tweet-text-container > p.tweet-text")
	selection.Each(func(index int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}