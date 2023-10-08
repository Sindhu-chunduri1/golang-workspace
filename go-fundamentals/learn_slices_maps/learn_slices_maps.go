package main

import (
	//"bytes"
	"fmt"
)

func main() {
	uniHello := "Hello world"
	bytes := []byte(uniHello)
	fmt.Println(bytes)
	runes := []rune(uniHello)
	fmt.Println(runes)
	runes[1] = 'a'
	fmt.Println(runes)
	fmt.Println(uniHello)
}
