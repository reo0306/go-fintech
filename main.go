package main

import (
	"fmt"
	"os/user"
	"time"
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

func main() {
	fmt.Println("Hello, World!", time.Now())
	fmt.Println(user.Current())

	fmt.Printf("Integer: %d\n", i)
	fmt.Printf("Float64: %f\n", f64)
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Boolean (true): %t\n", t)
	fmt.Printf("Boolean (false): %t\n", f)

	foo()
}