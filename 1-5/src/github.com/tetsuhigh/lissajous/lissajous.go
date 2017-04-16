package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// 先頭の色が初期背景色になる。
var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.White}

// 定数、パッケージ全体で見える
const (
	whiteIndex = 3 // 白のバレットでの位置
	blackIndex = 0 // 黒のバレットでの位置
	greenIndex = 1 // 緑のバレットでの位置
)

/*
  ランダムなリサージュ図形のGIFアニメーションを生成します。
 "lissajous > out.gif"とすることで、out.gifに出力することができる
*/
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	// 定数、関数内のみで見える
	const (
		cycles  = 5     // 発振器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンパスは[-size..+size]の範囲を扱う
		nframes = 64    //アニメーションフレーム数
		delay   = 8     // 10ms単位でのフレーム間の遅延
	)

	freq := rand.Float64() * 3.0 // 発振器yの相対周波
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		// 201x201のpaletteの先頭の色で塗りつぶされた画像を生成する
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// (x, y)の位置の色を変更する
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)	
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
