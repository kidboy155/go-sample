package main

import (
	"fmt"
	"time"
)

func TimeTrack1(start time.Time, functionName string) {
	elapesd := time.Since(start)
	fmt.Println(functionName, "took", elapesd)
}

func startPusher(queue chan<- int) chan<- bool {
	stopChan := make(chan bool)
	go func() {
		for i := 1; i <= 100; i++ {
			select {
			case <-stopChan:
				return
			default:
				time.Sleep(time.Millisecond * 10)
				queue <- i
			}
		}
	}()
	return stopChan
}
func main_crawl() {
	defer TimeTrack1(time.Now(), "Crawling: ")
	numberOfRequests := 1000
	maxWorkerNumber := 5
	queueChan := make(chan int, numberOfRequests)
	doneChan := make(chan int)

	for i := 1; i <= maxWorkerNumber; i++ {
		go func(name string) {
			for v := range queueChan {
				crawl(name, v)
			}
			fmt.Printf("%s is done\n", name)
			doneChan <- 1
		}(fmt.Sprintf("%d", i))
	}

	stopCh := startPusher(queueChan)
	stopCh1 := startPusher(queueChan)
	stopCh2 := startPusher(queueChan)

	go func() {
		time.Sleep(time.Second * 5)
		stopCh <- true
		stopCh1 <- true
		stopCh2 <- true

	}()
	// for i := 1; i <= numberOfRequests; i++ {
	// 	queueChan <- i
	// }
	// close(queueChan)
	// time.Sleep(time.Second * 5)
	for i := 1; i <= maxWorkerNumber; i++ {
		<-doneChan
	}
}

func crawl(name string, v int) {
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("Worker %s is crawling: %d \n", name, v)
}
