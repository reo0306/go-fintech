package main

import (
	"fmt"
)

func hy2(num int) string {
	if num %2 == 0 {
		return "ok"
	} else {
		return "ng"
	}
}

func main() {
	result := hy2(10)

	if result == "ok" {
		fmt.Println("The number is even")
	}

	if result2 := hy2(10); result2 == "ok" {
		fmt.Println("The number is even2")
	}

	num := 9
	if num %2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else if num %3 == 0 {
		fmt.Printf("%d is divisible by 3\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	x, y := 11, 12

	if x == 10 && y == 10 {
		fmt.Println("x and y are both 10")
	}

	if x == 10 || y == 10 {
		fmt.Println("Either x or y is 10")
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}

		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println("final sum:", sum)
}