package main

import (
	"fmt"
	"time"
	"regexp"
	"sort"
	"context"
	"io/ioutil"
	//"log"
	"bytes"
)

const  (
	c1 = iota
	c2
	c3
)

const (
	_ = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// 1回
    match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	// 複数
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	//s := "/view/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

	i := []int{5, 3, 2, 6, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Mike", 40},
		{"Vera", 45},
	}
	fmt.Println(i, s, p)

	sort.Ints(i)
	fmt.Println(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)

	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)

	ch := make(chan string)
	//ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, 0 * time.Second)
	//defer cancel()
	ctx := context.TODO()
	go longProcess(ctx, ch)
    //cancel()

	CTXLOOP:
	for {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <- ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}

	fmt.Println("###############")

	/*context2, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(context2))

	if err := ioutil.WriteFile("ioutil_temp.go", context2, 0666); err != nil {
		log.Fatal(err)
	}*/

	rr := bytes.NewBuffer([]byte("abc"))
	context2, _ := ioutil.ReadAll(rr)
	fmt.Println(string(context2))
}
