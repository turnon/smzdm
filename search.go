package smzdm

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	https = "https:"
	root  = "https://search.smzdm.com/"
	query = root + "/?v=b&c=home&s="
)

type entry struct {
	Title, Price, Time, Img, Href string
}

func (e *entry) extract(s *goquery.Selection) *entry {
	a := s.Find(".feed-block-title a").First()
	e.Href, _ = a.Attr("href")
	e.Title = strings.TrimSpace(a.Text())
	e.Price = strings.TrimSpace(s.Find(".feed-block-title a div").First().Text())
	timeBlock := s.Find(".feed-block-extras").First()
	timeBlock.Children().Remove()
	e.Time = strings.TrimSpace(timeBlock.Text())
	img, _ := s.Find("img").First().Attr("src")
	e.Img = https + img
	return e
}

type Search struct {
	doc     *goquery.Document
	Index   int
	Keyword string
	Entries []*entry
}

func Query(keyword string) *Search {
	s := Search{Keyword: keyword}
	s.ing()
	s.extract()
	return &s
}

func (s *Search) ing() *Search {
	key := url.QueryEscape(s.Keyword)
	resp, err := http.Get(query + key)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	s.doc, err = goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		panic(err)
	}

	return s
}

func (s *Search) extract() {
	s.doc.Find("#feed-main-list .feed-block").Each(func(i int, selection *goquery.Selection) {
		e := new(entry).extract(selection)
		s.Entries = append(s.Entries, e)
	})
}
