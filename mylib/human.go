package mylib

import "fmt"

var Public string = "Public"

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human")
}