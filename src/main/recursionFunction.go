package main

import "fmt"

func main() {
	i := 15
	fmt.Printf("%d's factorial is %d\n", i, factorial(i))
}

func factorial(i int) (result int) {
	if i == 0 {
		return 1
	} else {
		result = i * factorial(i-1)
	}
	return
}
