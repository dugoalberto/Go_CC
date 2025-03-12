package main

import (
	"fmt"
	"os"
)

/*
	func main() {
		var x = [100]int{}
		var k = 0
		for k < len(x) {
			x[k] = k
			k++
		}
		fmt.Println(x)
		k = 0
		for k < len(x) {
			if x[k]%3 == 0 {
				fmt.Println(x[k])
			}
			k++
		}
	}
*/
func main() {
	//albertodugo@MacBook-Air-di-Alberto myFirstProject % go build
	//albertodugo@MacBook-Air-di-Alberto myFirstProject % ./myFirstProject hello word 23
	//hello word 23
	//you need to build it and later on call it by
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
