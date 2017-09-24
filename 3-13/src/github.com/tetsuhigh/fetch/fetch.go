package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
 練習問題1-7 引数で指定されたURLにある内容を表示する
*/
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// ReadAllよりもCopyのほうがメモリに読み込む必要がないため効率的
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
