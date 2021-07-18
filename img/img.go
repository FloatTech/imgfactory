package img

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

// 获取高，宽
func Wh(i image.Image) (w, h int) {
	w = i.Bounds().Size().X
	h = i.Bounds().Size().Y
	return
}

func GuanZhu(path, qq, txt1, txt2 string) {
	// 加载图片
	im1, _ := gg.LoadImage(fmt.Sprintf("%v\\guanzhu\\gg.png", path))
	yuanpath := fmt.Sprintf("%v\\user\\%v\\yuan.png", path, qq)
	Size(
		yuanpath,
		yuanpath,
		155,
		155,
	)
	im2, _ := gg.LoadImage(yuanpath)

	dc := gg.NewContext(480, 253)
	dc.DrawImage(im2, 45, 45)
	dc.DrawImage(im1, 0, 0)
	//字体，大小，颜色，位置
	dc.LoadFontFace(fmt.Sprintf("%v\\font\\3.ttf", path), 38)
	dc.SetRGB255(0, 0, 0)
	dc.DrawString(txt1, 210, 110)
	// dc.Clear()
	dc.LoadFontFace(fmt.Sprintf("%v\\font\\3.ttf", path), 36)
	dc.SetRGB255(190, 190, 190)
	dc.DrawString(txt2, 210, 170)
	dc.SavePNG(fmt.Sprintf("%v\\user\\%v\\guanzhu.png", path, qq))
	// dc.SavePNG("guanzhu.png")
}

func Yuan(path, name string) {
	// 加载图片，这里路径换成自己的
	im, _ := gg.LoadImage(fmt.Sprintf("%v\\%v", path, name))
	// 获取图片的宽度和高度
	w, h := Wh(im)

	dc := gg.NewContext(w, h)

	// 取宽度和高度的最小值作为直径
	radius := math.Min(float64(w), float64(h)) / 2
	// 画圆形
	dc.DrawCircle(float64(w/2), float64(h/2), radius)
	// 对画布进行裁剪
	dc.Clip()
	// 加载图片
	dc.DrawImage(im, 0, 0)
	dc.SavePNG(fmt.Sprintf("%v\\yuan%v", path, name))
}

func DisYuan(path, name string, x, y, r float64) {
	// 加载图片，这里路径换成自己的
	im, _ := gg.LoadImage(fmt.Sprintf("%v\\%v", path, name))
	// 获取图片的宽度和高度
	w, h := Wh(im)

	dc := gg.NewContext(w, h)
	dc.DrawCircle(x, y, r)
	dc.Clip()
	dc.InvertMask()
	// dc.DrawRectangle(0, 0, float64(w), float64(h))
	dc.DrawImage(im, 0, 0)
	dc.SavePNG(fmt.Sprintf("%v\\disyuan%v", path, name))
}

func Pa(path, qq string) {
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(59) + 1
	im1, _ := gg.LoadImage(fmt.Sprintf("%v\\pa\\%v.png", path, rand))
	yuanpath := fmt.Sprintf("%v\\user\\%v\\yuan.png", path, qq)
	Size(
		yuanpath,
		yuanpath,
		100,
		100,
	)
	im2, _ := gg.LoadImage(yuanpath)
	dc := gg.NewContext(500, 500)
	dc.DrawImage(im2, 0, 400)
	dc.DrawImage(im1, 0, 0)
	dc.SavePNG(fmt.Sprintf("%v\\user\\%v\\pa.png", path, qq))

	// dc.SavePNG("pa.png")

}

//缩放图片
func Size(path1, path2 string, w, h int) {
	im, _ := gg.LoadImage(path1)
	im1 := imaging.Resize(im, w, h, imaging.Lanczos)
	dst := imaging.New(w, h, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, im1, image.Pt(0, 0))
	imaging.Save(dst, path2)
}
