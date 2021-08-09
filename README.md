# img

>需要另外下载素材包配合使用:https://github.com/tdf1939/sucai


#png
##处理中图像
>type Dc struct {
	Im *image.NRGBA
	W  int
	H  int
}

##加载图片
func Load(path string) image.Image {

##设置底图
func NewDc(w, h int, fillColor color.Color) *Dc {

##载入图片作底图
func ImDc(path string, w, h int) *Dc {

##变形
func Size(im image.Image, w, h int) *Dc {

##横向合并图片
func AndW(im []*image.NRGBA) *Dc {

##纵向合并图片
func AndH(im []*image.NRGBA) *Dc {

##文本框 字体，大小，颜色 ，背景色，文本
func Text(font string, size float64, col []int, col1 []int, txt string) *Dc {

##保存png
func SavePng(im image.Image, path string) {

##颜色
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

#Dc
##克隆
func (dst *Dc) Clone() *Dc {

##变形
func (dst *Dc) Size(w, h int) *Dc {

##上部插入图片
func (dst *Dc) Over(im image.Image, w, h, x, y int) *Dc {

##底部插入图片
func (dst *Dc) DstOver(im image.Image, w, h, x, y int) *Dc {

##获取圆图
func (dst *Dc) Circle(r int) *Dc {

##剪取方图
func (dst *Dc) Clip(w, h, x, y int) *Dc {

##裁取圆图
func (dst *Dc) ClipCircle(x, y, r int) *Dc {

##扣取圆
func (dst *Dc) DstClipCircle(x, y, r int) *Dc {

##插入文本
func (dst *Dc) Text(font string, size float64, col []int, x, y float64, txt string) *Dc {


##饱和度(-100, 100)
func (dst *Dc) AdjustSaturation(a float64) *Dc {

##锐化
func (dst *Dc) Sharpen(a float64) *Dc {

##模糊图像
func (dst *Dc) Blur(a float64) *Dc {

##灰度
func (dst *Dc) Grayscale() *Dc {

##反色
func (dst *Dc) Invert() *Dc {

##浮雕
func (dst *Dc) Convolve3x3() *Dc {


#gif
##使用以下函数将image.Image转换为 *Paletted。最多256色
func GetPaletted(im image.Image) *image.Paletted {

##合并成gif,1 dealy 10毫秒
func AndGif(delay int, im []*image.NRGBA) *gif.GIF {

##保存gif
func SaveGif(g *gif.GIF, path string) {
