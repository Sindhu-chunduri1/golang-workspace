package main

import (
	"fmt"
	"reflect"
)

type details struct {
	fname string
	lname string
	age   int
}

type myType string

func showDetails(i, j interface{}) {
	t1 := reflect.TypeOf(i)
	k1 := t1.Kind()
	t2 := reflect.TypeOf(j)
	k2 := t2.Kind()
	fmt.Println("Type of first interface:", t1)
	fmt.Println("Kind of first interface:", k1)
	fmt.Println("Type of second interface:", t2)
	fmt.Println("Kind of second interface:", k2)
}

func main() {
	iD := myType("CSS059")
	person := details{
		fname: "Sai",
		lname: "Sindhu",
		age:   22,
	}
	showDetails(person, iD)
}
