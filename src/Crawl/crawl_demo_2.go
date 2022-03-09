package main

import (
	"fmt"
	// "time"
	// "timetrack"
)

func crawlData(urls []int, ch chan int, cap int) {
	for i := 0; i < cap; i++ {
		ch <- urls[i]
	}
}

// Chỗ này hơi hard code em chưa biết sửa ntn
func printData(channels []chan int, nRequests int) {
	for cnt := 0; cnt < nRequests; cnt++ {
		select {
		case url1 := <-channels[0]:
			fmt.Println(url1)
		case url2 := <-channels[1]:
			fmt.Println(url2)
		case url3 := <-channels[2]:
			fmt.Println(url3)
		case url4 := <-channels[3]:
			fmt.Println(url4)
		case url5 := <-channels[4]:
			fmt.Println(url5)
		}
	}
}
// Ý tưởng là chia đều n requests ra n channels chạy song song. Thằng nào trả về trước thì in ra trước
func main_demo_2() {
	// defer TimeTrack(time.Now(), "Crawl: ")
	nRequests := 100000
	urls := make([]int, nRequests)

	nChannels := 5
	for i := 0; i < nRequests; i++ {
		urls[i] = i
	}

	var channels []chan int
	for i := 0; i < nChannels; i++ {
		balanceBufferCap := nRequests / nChannels
		ch := make(chan int, balanceBufferCap)
		channels = append(channels, ch)

		start := i * balanceBufferCap
		end := (i+1)*balanceBufferCap + 1
		subUrls := urls[start:end]
		go crawlData(subUrls, channels[i], cap(channels[i]))
	}

	printData(channels, nRequests)
}
