package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		log.Fatal("Error: only 1 argument needed. Found 0 or more than 1")
	}

	resp, err := http.Get(args[1])
	if err != nil {
		log.Fatal("Error when downloading a page")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("Error: returned status code is", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error: couldn't create document for parsing")
	}

	var links []string
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("src")
		if exists {
			links = append(links, link)
		}
	})

	for n, link := range links[8:] {
		f, err := os.Create(strconv.Itoa(n))
		if err != nil {
			log.Fatal("Error: couldn't create a file")
		}
		defer f.Close()

		resp, err := http.Get(link)
		if err != nil {
			log.Fatal("Error: couldn't download image")
		}
		defer resp.Body.Close()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			log.Fatal("Error: couldn't save an image")
		}
	}
}
