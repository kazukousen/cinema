package scrape

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// Run ...
func Run(base, path, selector string) error {
	// fetch from url
	url := fmt.Sprintf("%s%s", base, path)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	// fetch from selector
	links := getLinks(base, doc.Find("#list_b").Find("li > a"))
	for _, link := range links {
		fmt.Println(link)
		sel, err := link.getBody()
		if err != nil {
			continue
		}
		fmt.Println(sel.Html())
	}
	return nil
}

func html(sel *goquery.Selection) {
	fmt.Println(sel.Html())
}

func getLinks(base string, anchors *goquery.Selection) []Link {
	links := make([]Link, anchors.Size())
	anchors.Each(func(i int, s *goquery.Selection) {
		if url, ok := s.Attr("href"); ok {
			links[i] = Link{
				LinkBase: LinkBase{
					BaseURL: base,
				},
				URL:  url,
				Name: s.Text(),
			}
		}
	})
	return links
}
