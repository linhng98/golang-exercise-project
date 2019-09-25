package main

import (
	"errors"
	"fmt"
	"sync"
)

// Fetcher interface
type Fetcher interface {
	// Fetch func return body and slice of urls found on that page
	Fetch(url string) (body string, urls []string, err error)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var errLoading = errors.New("this url is loading")

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func (f *fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found %s", url)
}

func crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		fmt.Printf("done with %s. depth = 0\n", url)
	}

	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		fmt.Printf("vvvv %s has been crawled, not crawl again\n", url)
		return
	}
	// mark as loading avoid another rountine fetch this url
	fetched.m[url] = errLoading
	fetched.Unlock()

	_, urls, err := fetcher.Fetch(url)

	// update status in synced zone
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	// catch fetch error
	if err != nil {
		fmt.Printf("xxxx error occur on %s: %s\n", url, err)
		return
	}

	fmt.Printf("Found %s\n", url)
	done := make(chan bool)
	for i, u := range urls {
		fmt.Printf("---> crawling child %d/%d of %s.  %s\n", i+1, len(urls), url, u)
		go func(url string) {
			crawl(url, depth-1, fetcher)
			done <- true
		}(u)
	}

	for i := 0; i < len(urls); i++ {
		<-done
	}
}

func main() {
	crawl("https://golang.org/pkg/", 4, fetcher)

	fmt.Printf("\n\n")
	for u, err := range fetched.m {
		if err != nil {
			fmt.Printf("failed %s %s.\n", u, err)
		} else {
			fmt.Printf("%s was fetched.\n", u)
		}
	}
}
