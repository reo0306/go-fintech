package main

import (
	"fmt"
	"time"
	"regexp"
)

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
}
