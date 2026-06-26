package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port int
	DbName string
	SQLDriver string
}

var Config ConfigList

// go ルーチンで実行する数を制御できる
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}
	/*if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}*/
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList {
		Port: cfg.Section("web").Key("port").MustInt(),
		DbName: cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}
func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)

	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
