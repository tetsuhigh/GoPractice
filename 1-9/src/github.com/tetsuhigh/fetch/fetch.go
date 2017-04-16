package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
 練習問題1-9 引数で指定されたURLにある内容、HHTPステータスコードを表示する
 先頭にhttp://、もしくはhttps://がなければ追加する
*/
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "%s\n", resp.Status)
		// ReadAllよりもCopyのほうがメモリに読み込む必要がないため効率的
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
