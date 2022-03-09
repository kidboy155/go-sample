package timetrack


import (
	"fmt"
	"time"
	"github.com/fatih/color"
)


func TimeTrack(start time.Time, functionName string) {
	elapesd := time.Since(start)
	Println(functionName, "took", elapesd)
}

func Println(str string) {
	if len(str) == 0 {
		return
	}
	color.Red(str)
 }