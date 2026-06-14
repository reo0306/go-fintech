package main

import (
	"fmt"
	"sync"
	//"time"
)

func goroutine(s string, wg *sync.WaitGroup) {
    for i := 0; i < 5; i++ {
        //time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
    wg.Done()
}


func normal(s string) {
    for i := 0; i < 5; i++ {
        //time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func goroutine1(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}

/*func goroutine2(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}*/

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go goroutine("world", &wg)
    normal("hello")
    wg.Wait()

    s := []int{1, 2, 3, 4, 5}
    c := make(chan int) // 15, 15
    go goroutine1(s, c)
    go goroutine1(s, c)
    x := <-c
    fmt.Println("sum:", x)
    y := <-c
    fmt.Println("sum:", y)

    ch := make(chan int, 2)
    ch <- 100
    fmt.Println(len(ch))
    ch <- 200
    fmt.Println(len(ch))


    close(ch) // forの時は、closeが必要
    for c := range ch {
        fmt.Println(c)
    }
}
