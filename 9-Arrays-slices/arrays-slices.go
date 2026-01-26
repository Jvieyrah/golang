package main

import "fmt"

func main() {

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	var array3 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)

	slice1 = append(slice1, 6)
	fmt.Println(slice1)

	slice2 := array2[1:3]
	fmt.Println(slice2)
}
