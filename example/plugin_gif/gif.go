package plugin_gif

import (
	"fmt"
	"image"

	// "github.com/tdf1939/img"
	img "github.com/FloatTech/ZeroBot-Plugin/img"
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
