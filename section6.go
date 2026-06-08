package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

// 値れしバー
func (v Vertex) Area() int {
	return v.X * v.Y
}

//　ポインタれしバー
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	//fmt.Println(Area(v))
	v.Scale(10)
	fmt.Println(v.Area())
}
