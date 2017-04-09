package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 Dup2 入力に2回以上現れた行の数とその行のテキストを表示する。
 標準入力から読み込むか、名前が指定されたファイルの一覧から読み込む。
*/
func main() {
	// keyがstring、valueがintの空mapを生成する
	counts := make(map[string]int)
	files := os.Args[1:]
	
	if len(files) == 0 {
		fmt.Println("inout start")
		countLines(os.Stdin, counts)
	} else {
		// 第一戻り値はインデックスで"_"(アンダーバー)に代入することで破棄される
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				// 標準出力にエラー情報を出力
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	
	fmt.Println("dup result")
	// rangeはkey, valueを返す。順番は不定
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan() {
		// 空文字なら標準入力を終える。止め方わからない
		if input.Text() == "" {
			break
		}
		// 参照渡しなのでmainで渡されたオブジェクトを直接変更する
		counts[input.Text()]++
	}
}