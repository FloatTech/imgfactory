package plugin_gif

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	// "github.com/tdf1939/img"
	img "github.com/FloatTech/ZeroBot-Plugin/img"
)

//爬
func (cc Paths) Pa() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Circle(0).Im
	//随机爬图序号
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(60) + 1
	dc := img.ImDc(cc.Sc+`pa/`+strconv.Itoa(rand)+`.png`, 0, 0).DstOver(tou, 100, 100, 0, 400).Im
	img.SavePng(dc, cc.User+`爬.png`)
	return img.SGpic(cc.User + `爬.png`)

}

//撕
func (cc Paths) Si() string {
	tou := img.ImDc(cc.Pngs[0], 0, 0).Im
	im1 := img.Rotate(tou, 20, 380, 380)
	im2 := img.Rotate(tou, -12, 380, 380)
	dc := img.ImDc(cc.Sc+`si/0.png`, 0, 0).DstOver(im1.Im, im1.W, im1.H, -3, 370).DstOver(im2.Im, im2.W, im2.H, 653, 310).Im
	img.SavePng(dc, cc.User+`撕.png`)
	return img.SGpic(cc.User + `撕.png`)
}

//关注 必须下载素材字体才能使用
func (cc Paths) GuanZhu(txt1, txt2 string) string {
	// 加载图片
	im := img.ImDc(cc.Pngs[0], 0, 0)
	//叠加图片
	dst := img.ImDc(`data/image/sucai/vguanzhu/gg.png`, 0, 0).DstOver(im.Im, 155, 155, 45, 45)
	dc := dst.Text(fmt.Sprintf(`data/image/sucai/font/3.ttf`, cc.Sc), 38, img.Black, 210, 110, txt1).
		Text(fmt.Sprintf(`data/image/sucai/font/3.ttf`, cc.Sc), 36, img.Grey, 210, 170, txt2).
		Im
	img.SavePng(dc, cc.User+`关注.gif`)
	return img.SGpic(cc.User + `关注.gif`)

}
