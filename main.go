package main

import (
	"fmt"
	"os"

	"github.com/tdf1939/img/img"
)

func main() {
	// //加载图片
	// im := img.ImgPath("D:\\github\\sucai\\cs\\1.png")
	// //获取宽高
	// w, h := img.Wh(im)
	// fmt.Printf("%d %d\n", w, h)

	// 接收用户传递参数
	list := os.Args

	switch {
	case list[1] == "guanzhu":
		img.GuanZhu(list[2], list[3], list[4], list[5])
	case list[1] == "pa":
		img.Pa(list[2], list[3])
	default:
		fmt.Printf("空！\n")
	}
	// img.Pa("D:\\github\\sucai", "1073344346")
	// img.GuanZhu("D:\\github\\sucai", "1073344346", "哈哈哈", "hhhhhh")
}
