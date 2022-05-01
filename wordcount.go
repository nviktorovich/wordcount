package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//	'go is awesome'
	str := os.Args[1]
	cnt := 0
	if str == "" {
		fmt.Println(cnt)
	} else {
		for range strings.Split(str, " ") {
			cnt++
		}
		fmt.Println(cnt)
	}

}
