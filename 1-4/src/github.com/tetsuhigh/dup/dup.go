package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
 Dup4 入力に2回以上現れた行の数とその行のテキスト、含まれていたファイル名を表示する。
  名前が指定されたファイルの一覧から読み込む。
*/
func main() {
	// keyがstring、valueがintの空mapを生成する
	counts := make(map[string]int)
	files := make(map[string][]string)

	// 第一戻り値はインデックスで"_"(アンダーバー)に代入することで破棄される
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			// 標準出力にエラー情報を出力
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		
		// 改行文字が\r\nではないと一致判定されない
		for _, line := range strings.Split(string(data), "\r\n") {
			counts[line]++
			// appendの戻り値をsetする必要がある。appendだけでは追加されない。
			files[line] = append(files[line], filename)
		}
	}

	fmt.Println("dup result")
	// rangeはkey, valueを返す。順番は不定
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, files[line])
		}
	}
}
