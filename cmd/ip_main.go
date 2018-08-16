/*
用于测试 datx 是否正常释放, 无其他功能
*/
package main

import (
	"fmt"

	"github.com/molizz/goip"
)

func main() {
	goip.AddLocal()
	loc, err := goip.GetLocation("123.58.180.8")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(loc.ToString())
}
