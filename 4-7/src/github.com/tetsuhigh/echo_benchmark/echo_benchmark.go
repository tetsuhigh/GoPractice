package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

/*
 　練習問題1.3　1.1、1.2、strings.Joinを使ったechoのベンチマークを行う
*/
func main() {
	// 1.1の処理
	func1 := func() {
		fmt.Println(strings.Join(os.Args, " "))
	}
	
	// 1.2の処理
	func2 := func() {
		for i := 1; i < len(os.Args); i++ {
			fmt.Println(i, os.Args[i])
		}
	}

	// strings.Joinを使った処理
	func3 := func() {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}
	
	// printlnにフォーマットさせる処理
	func4 := func() {
		fmt.Println(os.Args[1:])
	}
	
	benchmark1 := benchmark(func1)
	benchmark2 := benchmark(func2)
	benchmark3 := benchmark(func3)
	benchmark4 := benchmark(func4)
	
	fmt.Println("----------------------------")
	fmt.Printf("1.1 %.2fs\n", benchmark1)
	fmt.Printf("1.2 %.2fs\n", benchmark2)
	fmt.Printf("strings.Join %.2fs\n", benchmark3)
	fmt.Printf("println %.2fs\n", benchmark4)
}

/*
　処理を実行し、実行時間をする
*/
func benchmark(function func()) float64 {
	start := time.Now()
	function()
	return time.Since(start).Seconds()
}