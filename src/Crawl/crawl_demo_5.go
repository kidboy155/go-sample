package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

func main() {
	stop := make(chan bool)
	defer close(stop)

	for item := range Crawl("https://golang.org/", 3, stop) {
		if item.error != nil {
			fmt.Printf("%v = %v\r\n", item.url, item.error.Error())
		} else {
			fmt.Printf("%v = %v\r\n", item.url, item.body)
		}
	}
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type FetchResult struct {
	url   string
	body  string
	error error
}

func Crawl(url string, concurrent int, stop <-chan bool) <-chan FetchResult {

	result := make(chan FetchResult)

	go func() {

		defer close(result)

		var counter int
		counterMutex := new(sync.Mutex)

		urlQueue := list.New()
		urlQueueMutex := new(sync.Mutex)

		visitedUrl := make(map[string]bool)
		visitedUrlMutex := new(sync.Mutex)

		urlQueue.PushBack(string(url))

		for {
			urlQueueMutex.Lock()
			node := urlQueue.Front()
			urlQueueMutex.Unlock()

			if node == nil {
				counterMutex.Lock()
				if counter > 0 {
					counterMutex.Unlock()
					continue
				} else {
					counterMutex.Unlock()
					urlQueueMutex.Lock()
					if urlQueue.Len() == 0 {
						urlQueueMutex.Unlock()
						// fmt.Println("No more task")
						break
					} else {
						urlQueueMutex.Unlock()
					}
				}
			} else {
				urlQueueMutex.Lock()
				urlQueue.Remove(node)
				urlQueueMutex.Unlock()

				counterMutex.Lock()
				if counter >= concurrent {
					counterMutex.Unlock()
					continue
				} else {
					counterMutex.Unlock()
				}
			}

			currentUrl := node.Value.(string)

			counterMutex.Lock()
			counter++
			counterMutex.Unlock()

			go func(currentUrl string, counter *int) {

				payload, urls, error := fetcher.Fetch(currentUrl)

				visitedUrlMutex.Lock()
				visitedUrl[currentUrl] = true
				visitedUrlMutex.Unlock()

				result <- FetchResult{
					body:  payload,
					error: error,
					url:   currentUrl,
				}

				if error == nil {
					for _, nestedUrl := range urls {

						visitedUrlMutex.Lock()
						_, ok := visitedUrl[nestedUrl]
						visitedUrlMutex.Unlock()

						if !ok {
							urlQueueMutex.Lock()
							urlQueue.PushBack(nestedUrl)
							urlQueueMutex.Unlock()
						}
					}
				}

				counterMutex.Lock()
				*counter--
				counterMutex.Unlock()
			}(currentUrl, &counter)
		}
	}()

	return result
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body    string
	elapsed time.Duration
	urls    []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		// time.Sleep(time.Duration(res.elapsed))
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found")
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		100 * time.Millisecond,
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
			"https://golang.org/internal/",
		},
	},
	"https://golang.org/internal/": &fakeResult{
		"Packages internal",
		100 * time.Millisecond,
		[]string{
			"https://golang.org/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		100 * time.Millisecond,
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
			"https://golang.org/pkg/container/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		100 * time.Millisecond,
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/container/": &fakeResult{
		"Package container",
		100 * time.Millisecond,
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/container/list",
			"https://golang.org/container/heap",
		},
	},
	"https://golang.org/pkg/container/list": &fakeResult{
		"Package list",
		1 * time.Millisecond,
		[]string{"https://golang.org/pkg/container/"},
	},
	"https://golang.org/pkg/container/heap": &fakeResult{
		"Package heap",
		1 * time.Millisecond,
		[]string{"https://golang.org/pkg/container/"},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		1 * time.Millisecond,
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
