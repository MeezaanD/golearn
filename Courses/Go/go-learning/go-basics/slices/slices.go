package slices

import "fmt"

func LearnSlices() {
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 15)
	fmt.Println(slice)
	slice = append(slice, 155)
	fmt.Println(slice)

	sliceWithSize := make([]int, 7)
	fmt.Println(sliceWithSize)

	for _, val := range slice {
		fmt.Println(val)
	}
}
