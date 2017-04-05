package main

import (
	"fmt"
	"os"
	"strings"
)

/*
 　練習問題1.1　プログラムを起動したコマンド名と引数を表示する
*/
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
