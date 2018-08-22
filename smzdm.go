package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type entry struct {
	search             *search
	time, title, price string
}

func newEntry(selection *goquery.Selection) *entry {
	title := strings.TrimSpace(selection.Find(".feed-block-title a").First().Text())
	price := strings.TrimSpace(selection.Find(".feed-block-title a div").First().Text())
	timeBlock := selection.Find(".feed-block-extras").First()
	timeBlock.Children().Remove()
	time := strings.TrimSpace(timeBlock.Text())
	return &entry{time: time, title: title, price: price}
}

func (e *entry) printf() {
	fmt.Printf(e.search.getFormatStr(), e.time, e.title, e.price)
}

type search struct {
	keyword                     string
	entries                     []*entry
	timeLen, titleLen, priceLen int
	formatStr                   string
}

func (s *search) process() {
	keyword := url.QueryEscape(s.keyword)

	resp, err := http.Get("http://search.smzdm.com/?c=home&v=b&s=" + keyword)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		panic(err)
	}

	s.printKeyword()

	doc.Find("#feed-main-list .z-feed-content").Each(func(i int, selection *goquery.Selection) {
		s.append(newEntry(selection))
	})

	for _, e := range s.entries {
		e.printf()
	}
}

func (s *search) append(e *entry) {
	timeLen := len(e.time)
	if s.timeLen < timeLen {
		s.timeLen = timeLen
	}

	titleLen := len(url.QueryEscape(e.title))
	if s.titleLen < titleLen {
		s.titleLen = titleLen
	}

	priceLen := len(url.QueryEscape(e.price))
	if s.priceLen < priceLen {
		s.priceLen = priceLen
	}

	e.search = s
	s.entries = append(s.entries, e)
}

func (s *search) getFormatStr() string {
	if s.formatStr == "" {
		s.formatStr = "%" + strconv.Itoa(s.timeLen) + "s | " +
			"%" + strconv.Itoa(s.titleLen) + "s | " +
			"%" + strconv.Itoa(s.priceLen) + "s\n"
	}
	return s.formatStr
}

func (s *search) printKeyword() {
	dash := strings.Repeat("-", (20 - len(s.keyword)))
	fmt.Println(s.keyword + " " + dash)
}

func main() {
	flag.Parse()
	keywords := flag.Args()

	if len(keywords) <= 0 {
		fmt.Println("no keyword given")
		return
	}

	for i, k := range keywords {
		if i > 0 {
			fmt.Println()
		}
		(&search{keyword: k}).process()
	}
}
