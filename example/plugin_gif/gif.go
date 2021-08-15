package plugin_gif

import (
	"fmt"
	"image"

	"github.com/tdf1939/img"
	// "bot/img" // 基础词库
)

//摸
func (cc Paths) Mo() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	mo := []*image.NRGBA{
		img.ImDc(fmt.Sprintf(`%v/mo/0.png`, cc.Sc), 0, 0).DstOver(tou, 80, 80, 32, 32).Im,
		img.ImDc(fmt.Sprintf(`%v/mo/1.png`, cc.Sc), 0, 0).DstOver(tou, 70, 90, 42, 22).Im,
		img.ImDc(fmt.Sprintf(`%v/mo/2.png`, cc.Sc), 0, 0).DstOver(tou, 75, 85, 37, 27).Im,
		img.ImDc(fmt.Sprintf(`%v/mo/3.png`, cc.Sc), 0, 0).DstOver(tou, 85, 75, 27, 37).Im,
		img.ImDc(fmt.Sprintf(`%v/mo/4.png`, cc.Sc), 0, 0).DstOver(tou, 90, 70, 22, 42).Im,
	}
	img.SaveGif(img.AndGif(1, mo), cc.User+`摸.gif`)
	return img.SGpic(cc.User + `摸.gif`)
}

//摸
func (cc Paths) Cuo() string {
	tou := img.ImDc(cc.Pngs[0], 110, 110).Circle(0).Im
	m1 := img.Rotate(tou, 72, 0, 0)
	m2 := img.Rotate(tou, 144, 0, 0)
	m3 := img.Rotate(tou, 216, 0, 0)
	m4 := img.Rotate(tou, 288, 0, 0)
	cuo := []*image.NRGBA{
		img.ImDc(fmt.Sprintf(`%v/cuo/0.png`, cc.Sc), 0, 0).DstOverC(tou, 0, 0, 75, 130).Im,
		img.ImDc(fmt.Sprintf(`%v/cuo/1.png`, cc.Sc), 0, 0).DstOverC(m1.Im, 0, 0, 75, 130).Im,
		img.ImDc(fmt.Sprintf(`%v/cuo/2.png`, cc.Sc), 0, 0).DstOverC(m2.Im, 0, 0, 75, 130).Im,
		img.ImDc(fmt.Sprintf(`%v/cuo/3.png`, cc.Sc), 0, 0).DstOverC(m3.Im, 0, 0, 75, 130).Im,
		img.ImDc(fmt.Sprintf(`%v/cuo/4.png`, cc.Sc), 0, 0).DstOverC(m4.Im, 0, 0, 75, 130).Im,
	}
	img.SaveGif(img.AndGif(5, cuo), cc.User+`搓.gif`)
	return img.SGpic(cc.User + `搓.gif`)
}

//冲
func (cc Paths) Chong() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	chong := []*image.NRGBA{
		img.ImDc(fmt.Sprintf(`%v/xqe/0.png`, cc.Sc), 0, 0).Over(tou, 30, 30, 15, 53).Im,
		img.ImDc(fmt.Sprintf(`%v/xqe/1.png`, cc.Sc), 0, 0).Over(tou, 30, 30, 40, 53).Im,
	}
	img.SaveGif(img.AndGif(1, chong), cc.User+`冲.gif`)
	return img.SGpic(cc.User + `冲.gif`)

}
