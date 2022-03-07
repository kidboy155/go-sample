package main

import (
	"fmt"
	"module/c4f"

	"github.com/leekchan/accounting"
)

/*
export GO111MODULE="on"
export GO111MODULE=“off”
go mod init module
go build
go get github.com/leekchan/accounting
go get github.com/fatih/color

go mod graph
*/
func main_module() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))

	c4f.Println("That test")
}
