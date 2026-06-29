package main

import (
	"fmt"
	"go-fintech/gotrading/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
