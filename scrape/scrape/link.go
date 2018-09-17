package scrape

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// LinkBase ...
type LinkBase struct {
	BaseURL string
}

// Link ...
type Link struct {
	LinkBase
	URL  string
	Name string
}

func (l Link) String() string {
	return fmt.Sprintf("url=%s, name=%s", l.URL, l.Name)
}

func (l Link) fullPath() string {
	return fmt.Sprintf("%s%s", l.BaseURL, l.URL)
}

func (l Link) getBody() (*goquery.Selection, error) {
	doc, err := goquery.NewDocument(l.fullPath())
	if err != nil {
		return nil, err
	}
	return doc.Find("body").Find("#flickList2"), nil
}
