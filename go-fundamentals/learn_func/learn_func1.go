package main

import "fmt"

func main() {
	myaddOne := func(a int) int {
		return a + 1
	}
	fmt.Println(myaddOne(1))
}
