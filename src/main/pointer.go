package main

import "fmt"

const MAX = 3

func main() {
	var i int = 9
	var p *int
	p = &i
	fmt.Printf("variable address is: %x\n", &i)
	fmt.Printf("the pointer saved address is: %x\n", p)
	fmt.Printf("ths pointer address's value is: %d\n", *p)
	// null pointer nil
	var n *int
	fmt.Printf("variable address's value is: %d\n", n)
	if n == nil {
		fmt.Printf("the pointer is nill")
	}
	pointerArray()
	pointerToPointer()
}

func pointerArray() {
	var pArray [MAX]*int
	a := []int{3, 4, 5}
	var i int

	for i = 0; i < MAX; i++ {
		pArray[i] = &a[i]
	}

	fmt.Println(pArray)

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *pArray[i])
	}
}

func pointerToPointer() {
	var a int = 3000
	var p *int = &a
	var pp **int = &p

	/* get pp's value */
	fmt.Printf("variable a = %d\n", a)
	fmt.Printf("pointer variable *p = %d\n", *p)
	fmt.Printf("pointer to pointer **p = %d\n", **pp)
	fmt.Printf("pointer's addrss p = %x\n", &p)
	fmt.Printf("pointer to pointer's address  pp= %x\n", pp)
}
