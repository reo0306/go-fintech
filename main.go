package main

import "fmt"

func init() {
	fmt.Println("init")
}

func bazz() {
	fmt.Println("bazz")
}

func main() {
	bazz()
	fmt.Println("Hello, World!", "Test, Test")
}