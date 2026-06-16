package main

import (
	"fmt"
	"log"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, err := quote.NewQuoteFromTiingo("spy", "2026-01-01", "2026-04-01", quote.Daily, "aaa.csv")
	if err != nil {
		log.Fatalf("データの取得に失敗しました: %v", err)
	}

	// 2. データが本当に空じゃないかチェックする
	if len(spy.Close) == 0 {
		log.Fatal("データが0件です。APIトークンが正しく設定されているか確認してください。")
	}

	fmt.Print(spy.CSV())

	// 3. データがあること。を確認してからRSIを計算する
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}