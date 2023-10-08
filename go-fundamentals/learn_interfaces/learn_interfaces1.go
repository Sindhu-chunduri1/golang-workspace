package main

import (
	"fmt"
)

type Human interface {

	// Methods
	Language() string
}

type Person struct {
	Name string
}

func (obj *Person) Init(name string) {

	obj.Name = name
	fmt.Println("Person name: ", name)

}

func (obj *Person) Language() string {

	fmt.Println("Some language")
	return "Any language"
}

type SuperPerson struct {
	Name string
}

func (obj *SuperPerson) SuperInit(name string) {

	obj.Name = name
	fmt.Println("SuperPerson name: ", name)

}

func main() {

	fmt.Println("Sindhu")
	//create a person object with his/her language
	//var inter Human

	obj := new(Person)
	obj.Init("Chunduri Sindhu")
	obj.Language()

	obj1 := new(Person)
	obj1.Init("Shomu")
	obj1.Language()

	obj2 := SuperPerson{}
	obj2.SuperInit("Super Chunduri Sindhu")

}
