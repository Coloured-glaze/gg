package main

import (
	"image"
	"strconv"

	"github.com/Coloured-glaze/gg"
	"golang.org/x/image/draw"
)

func main() {
	// 缩放图像
	dc := gg.NewContext(300, 300)
	src, err := gg.LoadImage("./gopher.png")
	if err != nil {
		panic(err)
	}
	dst := image.NewRGBA(image.Rect(0, 0, 300, 300)) // 使用矩形创建一个空图像
	for i := 0; i < 4; i++ {
		gg.Transformer(gg.ScaleStyle(i)).
			Scale(dst, image.Rect(0, 0, 200, 300),
				src, src.Bounds(), draw.Over, nil) // 缩放图像大小到 200 * 300
		dc.DrawImage(dst, 0, 0)
		dc.SavePNG("./gopher-" + strconv.Itoa(i) + ".png")
	}
}
