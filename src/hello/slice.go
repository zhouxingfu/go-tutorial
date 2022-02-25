package main


import "fmt"


func printSlice(a []int){
	fmt.Printf("%p\n", &a)
}

func printArray(a [2]int){
	fmt.Printf("%p\n", &a)
}

func slices(){

	a := [5]int{1, 2, 4, 5, 7}
	fmt.Printf("%p\n",&a)
	printSlice(a[0:2])
	//test(a)
}