package img

import (
	"image"
	"math"

	"github.com/fogleman/gg"
)

// 加载图片，path为图片路径
func ImgPath(path string) (im image.Image) {
	im, _ = gg.LoadImage(path)
	return
}

// 获取高，宽
func Wh(i image.Image) (w, h int) {
	w = i.Bounds().Size().X
	h = i.Bounds().Size().Y
	return
}

func Out() {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}

func Yuan() {
	im := ImgPath("cs\\1.png")

	// 获取图片的宽度和高度
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y

	dc := gg.NewContext(h, w)

	// 取宽度和高度的最小值作为直径
	radius := math.Min(float64(w), float64(h)) / 2
	// 画圆形
	dc.DrawCircle(float64(w/2), float64(h/2), radius)
	// 对画布进行裁剪
	dc.Clip()
	// 加载图片
	dc.DrawImage(im, 0, 0)
	dc.SavePNG("out.png")
}
