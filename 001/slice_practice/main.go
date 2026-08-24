package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// slice2 := slice[2:7]

	slice = append(slice, 11, 12, 13)

	slice = append(slice[:4], slice[5:]...)

	for i := range slice {
		slice[i] *= 2
	}

	fmt.Println("切片slice: ", slice)
	fmt.Println("切片slice容量: ", cap(slice))

}
