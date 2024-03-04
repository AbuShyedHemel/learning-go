package main

import "fmt"

func main() {
	//variable declaration
	// var name string = "Hemel"
	// var birthYear int = 1996
	//x := "Hellow" //type inferred
	// var a, b, c int = 12, 13, 14
	// var arrayIntegerSample = [3]int{1, 2, 3}
	// arrayStringSample := [2]string{"Ok", "Ok"}
	// arrayStringSample[1] = "Hemel"
	// arrayDefaultValues := []int{}
	// initializeSpecificElement := [5]int{4: 40}
	mySlice := []int{1, 2, 3, 4, 5}

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))

	// fmt.Println(a, b, c)
	// fmt.Println("Name: ", name)
	// fmt.Println("Birth-Date: ", birthYear)
	// //fmt.Println("Inferred ", x)
	// fmt.Println("Hello World!")
	// fmt.Println("Array Sample", arrayStringSample[1])
	// fmt.Println("Array Sample", arrayDefaultValues)
	// fmt.Println("Initialize Specific Element Sample", initializeSpecificElement)
	// fmt.Println("Length of an Array :", len(initializeSpecificElement))
	fmt.Println("Slice :", mySlice)
	fmt.Println(cap(mySlice))
}
