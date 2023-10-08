package main

import (
	"fmt"
)

func main() {
	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println("Array1: ", arr1)
	fmt.Println("Array2: ", arr2)
	fmt.Println("Array3: ", arr3)

	// Access array elements
	fmt.Println("1st element of Array2: ", arr2[0])

	// Change array elements
	arr2[2] = 50
	fmt.Println("After changing 3rd element of Array2", arr2)
}
