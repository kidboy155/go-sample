package main

import (
	"fmt"
	"time"
	// "timetrack"
)

func crawlUrl(w int, ch chan int) {
	for i:= range ch {
		fmt.Printf("Worker %v crawled url %v\n", w, i)
	}
}

func main_demo_1() {
	// defer TimeTrack(time.Now(), "Crawl: ")
	urlCh := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go crawlUrl(i, urlCh)
	}
	for j := 0; j < 100; j++ {
		time.Sleep( 100 * time.Millisecond)
		urlCh <- j
	}
	fmt.Println("Finished")
} 