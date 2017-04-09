package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("start")
	// keyがstring、valueがintの空mapを生成する
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan() {
		// 空文字なら標準入力を終える。止め方わからない
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}
	
	// rangeはkey, valueを返す。順番は不定
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
