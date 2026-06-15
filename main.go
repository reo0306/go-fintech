package main

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"
	"time"

	"go-fintech/mylib"
	"go-fintech/mylib/under"
)

var (
	i    int     = 1
	f64  float64 = 1.0
	s    string  = "Go is great!"
	t, f bool    = true, false
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
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var ss string = "123"
	ii, _ := strconv.Atoi(ss)
	fmt.Printf("%T %v %d\n", ii, ii, ii)

	hh := "Hello, World!"
	fmt.Println(string(hh[0]))

	var a [2]int
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

	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	m["new"] = 500
	fmt.Println(m)

	v, ok := m["apple"]
	fmt.Printf("Value: %d, Exists: %t\n", v, ok)

	v2, ok2 := m["notthing"]
	fmt.Printf("Value: %d, Exists: %t\n", v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	bytes := []byte{72, 73}
	fmt.Println(bytes)
	fmt.Println(string(bytes))

	result1, result2 := add(10, 20)
	fmt.Println(result1, result2)

	result3 := cal(100, 2)
	fmt.Println(result3)

	fff := func() {
		fmt.Println("This is an anonymous function.")
	}
	fff()

	func(x int) {
		fmt.Println("This is an anonymous function.", x)
	}(1)

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	cc1 := circleArea(3.14)
	fmt.Println(cc1(2))

	cc2 := circleArea(3)
	fmt.Println(cc2(2))

	foofoo()
	foofoo(1, 2)
	foofoo(1, 2, 3)

	s10 := []int{1, 2, 3}

	// 展開する
	foofoo(s10...)

	fff1 := 1.11
	fmt.Println(int(fff1))

	// 5,6

	en := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v\n", en, en)

	sss := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(sss))

	mylib.Say()

	under.Hello()

	person := mylib.Person{Name: "Mike", Age: 30}
	fmt.Println(person)

	fmt.Println(mylib.Public)
}

func foofoo(params ...int) {
	fmt.Println(len(params), params)

	for _, param := range params {
		fmt.Println(param)
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func cal(price int, item int) (result int) {
	result = price * item
	return
}

func converter(price int) float64 {
	return float64(price)
}

func add(x, y int) (int, int) {
	return x + y, x - y
}
