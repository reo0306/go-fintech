package main

import (
	"fmt"
	"sync"
	"time"
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
        c <- sum
    }
    close(c)
}

/*func goroutine2(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum
}*/

func producer(ch chan int, i int) {
    ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
    for i := range ch {
        func () {
            defer wg.Done()
            fmt.Println("processi", i * 1000)
        }()
    }
    fmt.Println("consumer done")
}

func producer2(first chan int) {
    defer close(first)
    for i := 0; i < 10; i++ {
        first <- i
    }
}

func multi2(first <-chan int, second chan<- int) {
    defer close(second)
    for i := range first {
        second <- i * 2
    }
}

func multi4(second <-chan int, third chan<- int) {
    defer close(third)
    for i := range second {
        third <- i * 4
    }
}

func goroutine4(ch chan string) {
    for {
        ch <- "packet from 1"
        time.Sleep(1 * time.Second)
    }
}

func goroutine5(ch chan int) {
    for {
        ch <- 100
        time.Sleep(1 * time.Second)
    }
}

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

    s2 := []int{1, 2, 3, 4, 5}
    c2 := make(chan int, len(s2))
    go goroutine1(s2, c2)

    for i := range c2 {
        fmt.Println(i)
    }

    var wg2 sync.WaitGroup
    ch2 := make(chan int)
   
    // Producer
    for i2 := 0; i2 < 10; i2++ {
        wg2.Add(1)
        go producer(ch2, i2)
    }

    // Consumer
    go consumer(ch2, &wg2)
    wg2.Wait()
    close(ch2)
    time.Sleep(2 * time.Second) // Ensure consumer finishes before main exits
    fmt.Println("main done")

    first := make(chan int) 
    second := make(chan int) 
    third := make(chan int) 

    go producer2(first)
    go multi2(first, second)
    go multi4(second, third)
    for result := range third {
        fmt.Println(result)
    }

    c3 := make(chan string)
    c4 := make(chan int)
    go goroutine4(c3)
    go goroutine5(c4)

    for {
        select {
        case msg1 := <-c3:
            fmt.Println("Received from goroutine4:", msg1)
        case msg2 := <-c4:
            fmt.Println("Received from goroutine5:", msg2)
        }
    }
}
