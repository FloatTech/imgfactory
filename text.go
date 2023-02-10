package imgfactory

import (
	"bufio"
	"image"
	"strings"

	"github.com/FloatTech/gg"
)

func renderText(canvas *gg.Context, text string, width, fontSize int) (txtPic image.Image, err error) {
	buff := make([]string, 0, 32)
	s := bufio.NewScanner(strings.NewReader(text))
	line := strings.Builder{}
	for s.Scan() {
		for _, v := range s.Text() {
			length, _ := canvas.MeasureString(line.String())
			if int(length) <= width {
				line.WriteRune(v)
			} else {
				buff = append(buff, line.String())
				line.Reset()
				line.WriteRune(v)
			}
		}
		buff = append(buff, line.String())
		line.Reset()
	}
	_, h := canvas.MeasureString("好")
	canvas = gg.NewContext(width+int(h*2+0.5), int(float64(len(buff)*3+1)/2*h+0.5))
	canvas.SetRGB(1, 1, 1)
	canvas.Clear()
	canvas.SetRGB(0, 0, 0)
	for i, v := range buff {
		if v != "" {
			canvas.DrawString(v, float64(width)*0.01, float64((i+1)*3)/2*h)
		}
	}
	return canvas.Image(), nil
}

// RenderText 文字转图片 width 是图片宽度
func RenderText(text, font string, width, fontSize int) (txtPic image.Image, err error) {
	canvas := gg.NewContext(width, fontSize) // fake
	if err = canvas.LoadFontFace(font, float64(fontSize)); err != nil {
		return
	}
	return renderText(canvas, text, width, fontSize)
}

// RenderTextWith 文字转图片 width 是图片宽度
func RenderTextWith(text string, font []byte, width, fontSize int) (txtPic image.Image, err error) {
	canvas := gg.NewContext(width, fontSize) // fake
	if err = canvas.ParseFontFace(font, float64(fontSize)); err != nil {
		return
	}
	return renderText(canvas, text, width, fontSize)
}
