package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func hy2(num int) string {
	if num%2 == 0 {
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
	if num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else if num%3 == 0 {
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

	l := []string{"a", "b", "c"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	switch os := getOsName(); os {
	case "windows":
		fmt.Println("Windows OS")
	case "linux":
		fmt.Println("Linux OS")
	case "mac":
		fmt.Println("Mac OS", os)
	default:
		fmt.Println("Unknown OS")
	}

	t := time.Now()
	fmt.Println("Current time:", t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	/*defer fmt.Println("This will be printed last")

	foo()

	fmt.Println("This will be printed first")*/

	/*
		fmt.Println("This will be printed second")
		defer fmt.Println("This will be printed last1")
		defer fmt.Println("This will be printed last2")
		defer fmt.Println("This will be printed last3")

		fmt.Println("This will be printed first")
	*/

	/*file, _ := os.Open("./main.go")

	defer file.Close()

	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))*/

	/*LoggingSettings("./app.log")
	_, err := os.Open("./test.go")
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}

	log.Println("This is a log message")
	log.Printf("Current time: %s", time.Now().Format(time.RFC3339))

	log.Fatalf("%T %v", "error", "error!")
	log.Fatalln("error!")*/

	/*file, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln("Error opening file:", err)
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error changing directory:", err)
	}*/

	save()
	fmt.Println("This will be printed after save()")

	ll := []int{100, 300, 23, 11, 23, 2, 3, 6, 4}
	for i := 0; i < len(ll); i++ {
		if ll[i] == 2 {
			fmt.Println(ll[i])
		}
	}

	mm := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum2 := 0
	for _, v := range mm {
		sum2 += v
	}
	fmt.Println("Sum of values in map:", sum2)
}

func thirdPartyLib() {
	panic("Unable to connect to database")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println("Recovered from panic:", s)
	}()
	thirdPartyLib()
}

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	mutiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(mutiLogFile)
}

func foo() {
	defer fmt.Println("defer in foo")
	fmt.Println("foo")
}

func getOsName() string {
	return "mac"
}
