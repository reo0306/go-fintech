package main

import (
	"fmt"
)

func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)
	fmt.Println(&*&n)

	/*
	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)
	*/

	var p *int = new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	/*
	var p2 *int
	fmt.Println(p2)
	*p2++
	fmt.Println(p2)
	*/

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	var p3 *int = new(int)
	fmt.Printf("%T\n", p3)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var st = new(struct{})
	fmt.Printf("%T\n", st)
}
