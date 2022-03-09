package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		// go fmt.Println(i)
		go func(j int) {
			fmt.Println(j)
		}(i)

	}

	time.Sleep((time.Second))
}
