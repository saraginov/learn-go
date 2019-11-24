package main

import "fmt"

/*
	Problem description

	In this exercise you'll use Go's concurrency features to parallelize a web crawler

	Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice

	Hint: You can keep a cache of the URLs that have been fetched on a map, but maps alone are not
	safe for concurrent use!
*/

// Fetcher interface implements Fetch method
type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses Fetcher to recursively crawl pages starting with url, to a maximum depth
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: fetch URLs in paraller
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	// initially depth is 4, thus this does not return
	if depth <= 0 {
		return
	}

	// on iteration 1 running Fetch(golang.org)
	// body, is string, urls is a slice of strings and err is an Error
	// fakeResult is the struct we are receiving from Fetch along with err

	body, urls, err := fetcher.Fetch(url)

	// fmt.Println(body)
	// fmt.Println(urls)
	// fmt.Println(err)

	// if we encountered error, means no more children urls
	if err != nil {
		fmt.Println('\n')
		fmt.Println(err)
		return
	}

	fmt.Printf("found %s %q\n", url, body)

	for _, u := range urls {
		/*
			adding go before Crawl recursive call makes each recursive call a goroutine
			and remember "A goroutine is a lightweight thread managed by the Go runtime."
			meaning each thread runs in parallel

			go is insufficient on its own as Crawl is only called once
		*/
		go Crawl(u, depth-1, fetcher)
	}

	return
}

func main() {
	/*
		We can store urlCache in main or in Crawl,
		given the circumstances I think it's better to store it in main
		as we could use it elsewhere should we need to
	*/

	// declare cache => TODO: which of the 2 lines below is better and why?
	// var urlCache []string
	// urlCache := make([]string, 0)

	// call Crawl on golang.org, 4 levels deep, and fetcher is an instance of fakeFetcher
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
