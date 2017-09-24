package main

import (
	"fmt"
	"os"
)

/*
 　練習問題1.2　個々の引数のインデックスと値の組を1行ごとに表示する
*/
func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])	
	}
}
