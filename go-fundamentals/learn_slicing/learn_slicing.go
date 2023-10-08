package main

import "fmt"

func main() {
	s := make([]string, 0) //create slices using the make() function
	fmt.Println("length of s", len(s))
	s = append(s, "hello")
	fmt.Println(s)
}
