package main

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func static(keywords []string) func(writer http.ResponseWriter, request *http.Request) {
	var (
		timer      <-chan time.Time
		m          sync.Mutex
		htmlOutput output
	)

	fetch := func() {
		m.Lock()
		defer m.Unlock()

		timer = time.After(60 * time.Second)
		htmlOutput = new(html)
		process(keywords, htmlOutput)
		log.Println("pre-fetched")
	}

	fetch()

	go func() {
		for {
			<-timer
			fetch()
		}
	}()

	return func(writer http.ResponseWriter, request *http.Request) {
		m.Lock()
		defer m.Unlock()
		htmlOutput.print(writer)
	}
}

func dyna(writer http.ResponseWriter, request *http.Request) {
	q := request.URL.Query()["q"]
	if len(q) == 0 {
		return
	}
	keys := strings.Split(q[0], " ")
	log.Println(request.RemoteAddr, keys)
	htmlOutput := new(html)
	process(keys, htmlOutput).print(writer)
}
