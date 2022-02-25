package main

import "fmt"

func main(){
	// fmt.Println("HelloWorld!")
	// values()
	// vars()
	// ifs()
	// fors()
	// arrays()

	// a := [4]int{1,2,3,4}
	// fmt.Println(a)

	// var b = make([]string, 3)
	// fmt.Println(b)

	// slices()
	

	// var a = [3][2]int{
	// 	[2]int{},
	// 	[2]int{},
	// }

	// fmt.Println(a)

	//b := make(map[string]int)

	n := map[string]int{"foo": 1, "bar": 2}

	fmt.Println(n)

	nums := []int{2, 3, 4}
    sum := 0
    for num := range nums {
		fmt.Println(num)
        sum += num
    }
	fmt.Println(sum)
}

