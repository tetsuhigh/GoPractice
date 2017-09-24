package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"strconv"
)

/*
 練習問題 1-10  引数で渡されたURLを非同期でGetし、その結果（所要時間、バイト数）を表示するとともに、Bodyをファイルに出力する
*/
func main() {
	fetchall(os.Args[1:], "1")
	fetchall(os.Args[1:], "2")
}

func fetchall(urls []string, destFolderName string){
	start := time.Now()
	ch := make(chan string) // 文字列型のチャネルを作成する
	for i, url := range urls {
		go fetch(url, destFolderName + "\\file" + strconv.Itoa(i), ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // chチャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

/*
 URLからGetし、その結果（所要時間、バイト数）をチャネルに返す
 非同期で実行可能
*/
func fetch(url string, destPath string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // chチャネルへ送信
		return
	}
	
	dest, err := os.Create(destPath)
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	nbytes, err := io.Copy(dest, resp.Body) // nbytesが未宣言のため:=でよく、errは代入扱いになる
	defer resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}