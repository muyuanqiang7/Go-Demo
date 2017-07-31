package main

import (
	"unsafe"
	"fmt"
)

var _context string = "context variable"
var text = "菜鸟教程"
var flag bool
var _number float32 = 3.0

var (
	count int
	exist bool
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

const (
	i = 1 << iota
	j
	k
	l
)

const Pi float32 = 3.1415926

func main() {
	/* this is my first go application */
	//number, str := 123, "text"
	//var username = "muyuanqiang"
	//fmt.Println("Hello World!")
	//_number = Pi * _number * _number
	//println(_number)
	//println(a, b, c)
	//println(i, j, k, l)
	//println(_context, _number, flag, text, count, exist, number, str, username)
	//println(getMax(i, j))
	m, n := swap("kernel", "ubuntu")
	fmt.Println(m, n)
	println(getMax(5, 7))
}

func getMax(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}
