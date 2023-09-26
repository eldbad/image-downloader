package parse

import (
	"fmt"
	"io"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type HtmlParser struct{}

func (p HtmlParser) GetImageLinks(body io.ReadCloser) ([]*url.URL, error) {
	doc, err := goquery.NewDocumentFromReader(body)

	var links []*url.URL
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("src")
		if exists {
			link, err := url.Parse(link)
			if err != nil {
				fmt.Println("log, that there is err")
			}
			links = append(links, link)
		}
	})

	return links, err
}
