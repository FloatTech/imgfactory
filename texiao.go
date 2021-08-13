package img

import (
	"image/color"

	"github.com/disintegration/imaging"
)

//亮度(-100, 100)
func (dst *Dc) AdjustBrightness(a float64) *Dc {
	return &Dc{
		Im: imaging.AdjustBrightness(dst.Im, a),
		W:  dst.W,
		H:  dst.H,
	}
}

//对比度(-100, 100)
func (dst *Dc) AdjustContrast(a float64) *Dc {
	return &Dc{
		Im: imaging.AdjustContrast(dst.Im, a),
		W:  dst.W,
		H:  dst.H,
	}
}

//饱和度(-100, 100)
func (dst *Dc) AdjustSaturation(a float64) *Dc {
	return &Dc{
		Im: imaging.AdjustSaturation(dst.Im, a),
		W:  dst.W,
		H:  dst.H,
	}
}

//锐化
func (dst *Dc) Sharpen(a float64) *Dc {
	return &Dc{
		Im: imaging.Sharpen(dst.Im, a),
		W:  dst.W,
		H:  dst.H,
	}
}

//模糊图像 正数
func (dst *Dc) Blur(a float64) *Dc {
	return &Dc{
		Im: imaging.Blur(dst.Im, a),
		W:  dst.W,
		H:  dst.H,
	}
}

//灰度
func (dst *Dc) Grayscale() *Dc {
	b := dst.Im.Bounds()
	for y1 := b.Min.Y; y1 <= b.Max.Y; y1++ {
		for x1 := b.Min.X; x1 <= b.Max.X; x1++ {
			a := dst.Im.At(x1, y1)
			c := color.NRGBAModel.Convert(a).(color.NRGBA)
			f := 0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B) + 0.5
			c.R = uint8(f)
			c.G = uint8(f)
			c.B = uint8(f)
			dst.Im.Set(x1, y1, c)
		}
	}
	return dst
}

//反色
func (dst *Dc) Invert() *Dc {
	b := dst.Im.Bounds()
	for y1 := b.Min.Y; y1 <= b.Max.Y; y1++ {
		for x1 := b.Min.X; x1 <= b.Max.X; x1++ {
			a := dst.Im.At(x1, y1)
			c := color.NRGBAModel.Convert(a).(color.NRGBA)
			c.R = 255 - c.R
			c.G = 255 - c.G
			c.B = 255 - c.B
			dst.Im.Set(x1, y1, c)
		}
	}
	return dst
}

//浮雕
func (dst *Dc) Convolve3x3() *Dc {
	return &Dc{
		Im: imaging.Convolve3x3(
			dst.Im,
			[9]float64{
				-1, -1, 0,
				-1, 1, 1,
				0, 1, 1,
			},
			nil,
		),
		W: dst.W,
		H: dst.H,
	}
}
