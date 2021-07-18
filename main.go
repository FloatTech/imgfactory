package main

import (
	"fmt"

	"github.com/tdf1939/img/img"
)

func main() {
	//加载图片
	im := img.ImgPath("D:\\github\\sucai\\cs\\1.png")
	//获取宽高
	w, h := img.Wh(im)
	fmt.Printf("%d %d\n", w, h)

}
