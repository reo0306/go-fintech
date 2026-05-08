package main

import (
	"fmt"
	"os/user"
	"time"
	"strings"
	"strconv"
)

var (
	i int = 1
	f64 float64 = 1.0
	s string = "Go is great!"
	t, f bool = true, false
)

func foo() {
	xi := 1
	xi = 2
	var xf32 float32 = 2.0
	xs := "Go is awesome!"
	xt, xf := true, false
	
	fmt.Printf("Integer (short declaration): %d\n", xi)
	fmt.Printf("Float64 (short declaration): %f\n", xf32)
	fmt.Printf("String (short declaration): %s\n", xs)
	fmt.Printf("Boolean (true, short declaration): %t\n", xt)
	fmt.Printf("Boolean (false, short declaration): %t\n", xf)

	fmt.Printf("Type of xf32: %T\n", xf32)
	fmt.Printf("Type of xi: %T\n", xi)
}

const Pi = 3.14
const (
	Username = "admin"
	Password = "password123"
)

func main() {
	fmt.Println("Hello, World!", time.Now())
	fmt.Println(user.Current())

	fmt.Printf("Integer: %d\n", i)
	fmt.Printf("Float64: %f\n", f64)
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Boolean (true): %t\n", t)
	fmt.Printf("Boolean (false): %t\n", f)

	foo()

	fmt.Println(Pi, Username, Password)

	fmt.Println("Hello " + "World!")
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World!"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(`Test
	         TEST
		tesT`)

	t, f := true, false
	fmt.Printf("%T %v %t\n", t, 1, t)
	fmt.Printf("%T %v %t\n", f, 0, f)

	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy:= int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var ss string = "123"
	ii, _ := strconv.Atoi(ss)
	fmt.Printf("%T %v %d\n", ii, ii, ii)

	hh := "Hello, World!"
	fmt.Println(string(hh[0]))

	var a[2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)

	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)

	n = append(n, 100, 200, 300)
	fmt.Println(n)

	nn := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn)
	nn = append(nn, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn)
	nn = append(nn, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(nn), cap(nn), nn)

	aaa := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(aaa), cap(aaa), aaa)

	bb := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(bb), cap(bb), bb)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	//c = make([]int, 5) 
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}