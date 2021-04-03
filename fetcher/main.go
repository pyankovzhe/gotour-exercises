package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Save cache for fetched urls
type SafeCache struct {
	mu sync.RWMutex
	v  map[string]string
}

func (c *SafeCache) Insert(url string, body string) {
	c.mu.Lock()
	c.v[url] = body
	c.mu.Unlock()
}

func (c *SafeCache) Get(url string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if v, ok := c.v[url]; ok {
		return v, ok
	} else {
		return "", false
	}
}

var Cache = &SafeCache{v: make(map[string]string)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go crawlUrl(url, depth, fetcher, wg)
	wg.Wait()
}

func crawlUrl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	fmt.Printf("Processing url %s\n", url)

	if v, ok := Cache.Get(url); ok {
		fmt.Printf("found cached url: %s %q\n", url, v)
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	Cache.Insert(url, body)

	for _, u := range urls {
		wg.Add(1)
		go crawlUrl(u, depth-1, fetcher, wg)
	}

	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
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
