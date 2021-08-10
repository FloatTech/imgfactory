package img

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

//克隆
func (dst *Dc) Clone() *Dc {
	var src Dc
	src.Im = image.NewNRGBA(image.Rect(0, 0, dst.W, dst.H))
	draw.Over.Draw(src.Im, src.Im.Bounds(), dst.Im, dst.Im.Bounds().Min)
	src.W = dst.W
	src.H = dst.H
	return &src
}

//变形
func (dst *Dc) Size(w, h int) *Dc {
	dst = Size(dst.Im, w, h)
	return dst
}

//水平翻转
func (dst *Dc) FlipH() *Dc {
	return &Dc{
		Im: imaging.FlipH(dst.Im),
		W:  dst.W,
		H:  dst.H,
	}
}

//垂直翻转
func (dst *Dc) FlipV() *Dc {
	return &Dc{
		Im: imaging.FlipV(dst.Im),
		W:  dst.W,
		H:  dst.H,
	}
}

//上部插入图片
func (dst *Dc) Over(im image.Image, w, h, x, y int) *Dc {
	im1 := Size(im, w, h).Im
	//叠加图片
	draw.Over.Draw(dst.Im, dst.Im.Bounds(), im1, im1.Bounds().Min.Sub(image.Pt(x, y)))
	return dst
}

//底部插入图片
func (dst *Dc) DstOver(im image.Image, w, h, x, y int) *Dc {
	im1 := Size(im, w, h).Im
	dc := dst.Clone()
	dst = NewDc(dst.W, dst.H, color.NRGBA{0, 0, 0, 0})
	draw.Over.Draw(dst.Im, dst.Im.Bounds(), im1, im1.Bounds().Min.Sub(image.Pt(x, y)))
	draw.Over.Draw(dst.Im, dst.Im.Bounds(), dc.Im, dc.Im.Bounds().Min)
	return dst
}

//获取圆图
func (dst *Dc) Circle(r int) *Dc {
	if r == 0 {
		r = dst.H / 2
	}
	dst = dst.Size(2*r, 2*r)
	b := dst.Im.Bounds()
	for y1 := b.Min.Y; y1 < b.Max.Y; y1++ {
		for x1 := b.Min.X; x1 < b.Max.X; x1++ {
			if (x1-r)*(x1-r)+(y1-r)*(y1-r) > r*r {
				dst.Im.Set(x1, y1, color.NRGBA{0, 0, 0, 0})
			}
		}
	}
	return dst
}

//剪取方图
func (dst *Dc) Clip(w, h, x, y int) *Dc {
	dst.Im = dst.Im.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA)
	dst.W = w
	dst.H = h
	return dst
}

//裁取圆图
func (dst *Dc) ClipCircle(x, y, r int) *Dc {
	dst = dst.Clip(2*r, 2*r, x-r, y-r)
	b := dst.Im.Bounds()
	for y1 := b.Min.Y; y1 < b.Max.Y; y1++ {
		for x1 := b.Min.X; x1 < b.Max.X; x1++ {
			if (x1-x)*(x1-x)+(y1-y)*(y1-y) > r*r {
				dst.Im.Set(x1, y1, color.NRGBA{0, 0, 0, 0})
			}
		}
	}
	return dst
}

//扣取圆
func (dst *Dc) DstClipCircle(x, y, r int) *Dc {
	// dc := dst.Clip(x-r, y-r, 2*r, 2*r)
	b := dst.Im.Bounds()
	for y1 := b.Min.Y; y1 < b.Max.Y; y1++ {
		for x1 := b.Min.X; x1 < b.Max.X; x1++ {
			if (x1-x)*(x1-x)+(y1-y)*(y1-y) <= r*r {
				dst.Im.Set(x1, y1, color.NRGBA{0, 0, 0, 0})
			}
		}
	}
	return dst
}

//插入文本
func (dst *Dc) Text(font string, size float64, col []int, x, y float64, txt string) *Dc {

	dc := gg.NewContextForImage(dst.Im)
	//字体，大小，颜色，位置
	dc.LoadFontFace(font, size)
	dc.SetRGBA255(col[0], col[1], col[2], col[3])
	dc.DrawString(txt, x, y)
	ds := dc.Image()
	draw.Over.Draw(dst.Im, dst.Im.Bounds(), ds, ds.Bounds().Min)
	return dst
}
