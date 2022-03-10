package main

import (
	"fmt"
	"time"
)

func TimeTrack1(start time.Time, functionName string) {
	elapesd := time.Since(start)
	fmt.Println(functionName, "took", elapesd)
}
func main_demo_good() {
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

	for i := 1; i <= numberOfRequests; i++ {
		queueChan <- i
	}
	close(queueChan)
	// time.Sleep(time.Second * 5)
	for i := 1; i <= maxWorkerNumber; i++ {
		<- doneChan
	}
}

func crawl(name string, v int) {
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("Worker %s is crawling: %d \n", name, v)
}
