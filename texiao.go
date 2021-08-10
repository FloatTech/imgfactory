package img

import "github.com/disintegration/imaging"

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
	return &Dc{
		Im: imaging.Grayscale(dst.Im),
		W:  dst.W,
		H:  dst.H,
	}
}

//反色
func (dst *Dc) Invert() *Dc {
	return &Dc{
		Im: imaging.Invert(dst.Im),
		W:  dst.W,
		H:  dst.H,
	}
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
