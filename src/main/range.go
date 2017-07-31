package main

import "fmt"

func main() {
	var numbers = []int{2, 3, 4}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("the slice sum is: %d\n", sum)

	for i, value := range numbers {
		if value == 3 {
			fmt.Printf("this value %d's index is %d\n", value, i)
		}
	}

	kvs := map[string]string{"name": "muyuanqiang", "age": "28"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}
	//枚举字符串第一个参数index 第二个参数Unicode值字符
	for i, v := range "muyuanqiang" {
		fmt.Println(i, v)
	}
}
