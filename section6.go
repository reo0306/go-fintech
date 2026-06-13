package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

// 値レシバー
func (v Vertex) Area() int {
	return v.x * v.y
}

//　ポインタレシバー
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

//func Area(v Vertex) int {
	//return v.x * v.y
//}

type Vertex3D struct {
	Vertex
	z int
}

// 値レシバー
func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

//　ポインタレシバー
func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

type MyInt int

func (i MyInt) Double() int {
	return int(i * 2)
}

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr. " + p.Name
	fmt.Printf("My name is %s\n", p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr. Mike" {
		fmt.Println("Driving car...")
	} else {
		fmt.Println("You can't drive a car.")
	}
}

func do(i interface{}) {
	/*ii := i.(int) // タイプアサーション
	ii *= 2
	fmt.Printf("The result is %v\n", ii)

	ss := i.(string) // タイプアサーション
	fmt.Printf("The result is %v\n", ss + " world")
	*/
	switch v := i.(type) { // switch type文
	case int:
		fmt.Printf("The result is %v\n", v * 2)
	case string:
		fmt.Printf("The result is %v\n", v + " world")
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	//v := Vertex{3, 4}
	v := New(3, 4, 5)
	//fmt.Println(Area(v))
	v.Scale3D(10)
	//fmt.Println(v.Area())
	fmt.Println(v.Area3D())

	myInt := MyInt(10)
	fmt.Println(myInt.Double())

    var mike Human = &Person{"Mike"}
    var x Human = &Person{"X"}
    //var dog Dog = Dog{"Buddy"}
	//mike.Say()
	DriveCar(mike)
	DriveCar(x)
	//DriveCar(dog)

	//var i interface{} = 10
	do(10)
	do("hello")
	do(true)

	var i int = 10
	ii := float64(10)
	fmt.Printf("The type of ii is %T\n", ii)
}
