package main

import "fmt"

func main() {
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("the mean's value is:%f\n", mean)
}
