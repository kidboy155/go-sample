package main

import (
	"fmt"
	"time"
	// "timetrack"
)

func createList() []int {
	const MaxItem = 100
	inputList := make([]int, MaxItem)

	for i := 0; i < MaxItem; i++ {
		inputList[i] = i + 1
	}

	return inputList
}

func main_demo_3() {
	// defer TimeTrack(time.Now(), "Crawl: ")
	const MaxGoroutine = 5
	ch := make(chan int, MaxGoroutine)
	inputList := createList()

	go func(ch chan int, inputList []int) {
		for i := 0; i < len(inputList); i++ {
			ch <- inputList[i]
			fmt.Println("Process item", inputList[i])
		}
		defer close(ch)
	}(ch, inputList)

	for v := range ch {
		fmt.Println("read value", v)
		time.Sleep(time.Microsecond)
	}

	fmt.Println("Finish .....")
}
