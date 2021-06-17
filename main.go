package main

import (
	"flag"
	"net/http"
	"os"
	"sync"

	"github.com/turnon/smzdm/smzdm"
)

func main() {
	var result output
	toHTML := flag.Bool("h", false, "output html")
	webServer := flag.Bool("w", false, "work as webserver")
	flag.Parse()
	keywords := flag.Args()

	if *webServer {
		runServer(keywords)
		return
	}

	// keywords = append(keywords, "ascis")
	if len(keywords) <= 0 {
		panic("no keyword given")
	}

	if *toHTML {
		result = new(html)
	} else {
		result = new(stdout)
	}

	process(keywords, result).print()
}

func process(keywords []string, result output) output {
	searched := make(chan *smzdm.Search)
	var searchWg sync.WaitGroup
	searchWg.Add(len(keywords))

	go func() {
		for s := range searched {
			result.collect(s)
			searchWg.Done()
		}
	}()

	for i, k := range keywords {
		go func(k string, i int) {
			s := smzdm.Query(k)
			s.Index = i
			searched <- s
		}(k, i)
	}

	searchWg.Wait()
	close(searched)
	result.sort()
	return result
}

func runServer(specificKeys ...[]string) {
	var handler func(writer http.ResponseWriter, request *http.Request)
	keys := specificKeys[0]
	if len(keys) > 0 {
		handler = static(keys)
	} else {
		handler = dyna
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port(), nil)
}

func port() string {
	if port := os.Getenv("PORT"); len(port) != 0 {
		return port
	}
	return "80"
}
