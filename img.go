package img

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

//处理中图像
type Dc struct {
	Im *image.NRGBA
	W  int
	H  int
}

// 加载图片
func Load(path string) image.Image {
	var img image.Image
	if strings.HasPrefix(path, "http") {
		res, err := http.Get(path)
		if err != nil {
			fmt.Println("A error occurred!")
		}
		defer res.Body.Close()
		// 获得get请求响应的reader对象
		reader := bufio.NewReaderSize(res.Body, 32*1024)
		//读取路径
		img, _, _ = image.Decode(reader)
	} else {
		// 加载路径
		file, _ := os.Open(path)
		defer file.Close()
		//读取路径
		img, _, _ = image.Decode(file)
	}
	return img
}

//设置底图
func NewDc(w, h int, fillColor color.Color) *Dc {
	var dst Dc
	dst.W = w
	dst.H = h
	c := color.NRGBAModel.Convert(fillColor).(color.NRGBA)
	if (c == color.NRGBA{0, 0, 0, 0}) {
		dst.Im = image.NewNRGBA(image.Rect(0, 0, w, h))
	} else {
		dst.Im = &image.NRGBA{
			Pix:    bytes.Repeat([]byte{c.R, c.G, c.B, c.A}, w*h),
			Stride: 4 * w,
			Rect:   image.Rect(0, 0, w, h),
		}
	}
	return &dst
}

//载入图片作底图
func ImDc(path string, w, h int) *Dc {
	return Size(Load(path), w, h)
}

//变形
func Size(im image.Image, w, h int) *Dc {
	var dc Dc
	//修改尺寸
	if w > 0 && h > 0 {
		dc.W = w
		dc.H = h
		dc.Im = imaging.Resize(im, w, h, imaging.Lanczos)
	} else if w <= 0 && h <= 0 {
		dc.W = im.Bounds().Size().X
		dc.H = im.Bounds().Size().Y
		dc.Im = image.NewNRGBA(image.Rect(0, 0, dc.W, dc.H))
		draw.Over.Draw(dc.Im, dc.Im.Bounds(), im, im.Bounds().Min)
	} else if w == 0 {
		dc.H = h
		dc.W = h * im.Bounds().Size().X / im.Bounds().Size().Y
		dc.Im = imaging.Resize(im, dc.W, h, imaging.Lanczos)
	} else if h == 0 {
		dc.W = w
		dc.H = w * im.Bounds().Size().Y / im.Bounds().Size().X
		dc.Im = imaging.Resize(im, w, dc.H, imaging.Lanczos)
	}
	return &dc
}

//旋转
func Rotate(img image.Image, angle float64, w, h int) *Dc {
	im := Size(img, w, h)
	var dc Dc
	dc.Im = imaging.Rotate(im.Im, angle, color.NRGBA{0, 0, 0, 0})
	dc.W = dc.Im.Bounds().Size().X
	dc.H = dc.Im.Bounds().Size().Y
	return &dc
}

//横向合并图片
func AndW(im []*image.NRGBA) *Dc {
	dc := make([]*Dc, len(im))
	h := im[0].Bounds().Size().Y
	w := 0
	for i, value := range im {
		dc[i] = Size(value, 0, h)
		w += dc[i].W
	}
	ds := NewDc(w, h, color.NRGBA{0, 0, 0, 0})
	x := 0
	for _, value := range dc {
		ds = ds.Over(value.Im, value.W, h, x, 0)
		x += value.W
	}
	return ds
}

//纵向合并图片
func AndH(im []*image.NRGBA) *Dc {
	dc := make([]*Dc, len(im))
	w := im[0].Bounds().Size().X
	h := 0
	for i, value := range im {
		dc[i] = Size(value, 0, w)
		h += dc[i].H
	}
	ds := NewDc(w, h, color.NRGBA{0, 0, 0, 0})
	y := 0
	for _, value := range dc {
		ds = ds.Over(value.Im, w, value.H, 0, y)
		y += value.H
	}
	return ds
}

//文本框 字体，大小，颜色 ，背景色，文本
func Text(font string, size float64, col []int, col1 []int, txt string) *Dc {
	var dst Dc
	dc := gg.NewContext(10, 10)
	dc.SetRGBA255(0, 0, 0, 0)
	dc.Clear()
	dc.SetRGBA255(col[0], col[1], col[2], col[3])
	dc.LoadFontFace(font, size+size/2)
	w, h := dc.MeasureString(txt)
	w = w - size*2
	dc1 := gg.NewContext(int(w), int(h))
	dc1.SetRGBA255(col1[0], col1[1], col1[2], col1[3])
	dc1.Clear()
	dc1.SetRGBA255(col[0], col[1], col[2], col[3])
	dc1.LoadFontFace(font, size)
	dc1.DrawStringAnchored(txt, w/2, h/2, 0.5, 0.5)
	dst.Im = image.NewNRGBA(image.Rect(0, 0, int(w), int(h)))
	draw.Over.Draw(dst.Im, dst.Im.Bounds(), dc1.Image(), dc1.Image().Bounds().Min)
	dst.W, dst.H = dst.Im.Bounds().Size().X, dst.Im.Bounds().Size().Y
	return &dst
}

//保存png
func SavePng(im image.Image, path string) {
	f, _ := os.Create(path) //创建文件
	defer f.Close()         //关闭文件
	png.Encode(f, im)       //写入
}

//颜色
var White = []int{255, 255, 255, 255}
var Black = []int{0, 0, 0, 255}
var Red = []int{255, 0, 0, 255}
var Green = []int{0, 255, 0, 255}
var Blue = []int{0, 0, 255, 255}
var Yellow = []int{255, 255, 0, 255}
var Cyan = []int{0, 255, 255, 255}
var Magenta = []int{255, 0, 255, 255}
var Grey = []int{190, 190, 190, 255}
var Pink = []int{255, 181, 197, 255}
var Orange = []int{255, 165, 0, 255}
var TouM = []int{0, 0, 0, 0}
