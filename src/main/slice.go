package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	printSlice(numbers)
	var _numbers []int
	printSlice(_numbers)
	if _numbers == nil {
		fmt.Printf("slice is nill\n")
	}

	_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(_slice)

	fmt.Println("_slice ==", _slice)
	// cut slice
	fmt.Println("_slice[1:4]==", _slice[1: 4])
	fmt.Println("_slice[:3]==", _slice[: 3])
	fmt.Println("_slice[4:]==", _slice[4:])
	fmt.Println("---------------------- slice append and copy ------------------------")
	sliceAppendAndCopy()
}

func printSlice(numbers []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
}

func sliceAppendAndCopy() {
	var numbers []int
	printSlice(numbers)
	// append one value
	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	_numbers := make([]int, len(numbers), cap(numbers)*2)
	copy(_numbers, numbers)
	var i int
	for i = 0; i < 5; i++{
		_numbers =append(_numbers, i)
	}
	printSlice(_numbers)
}
