package img

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path"
)

//上传搜狗图床
func SGpic(s string) string {
	var buff bytes.Buffer
	// 创建一个Writer
	writer := multipart.NewWriter(&buff)
	// 写入一般的表单字段
	writer.WriteField("key", "value")
	// 写入图片字段
	// CreateFormFile第一个参数是 表单对应的字段名
	// 第二个字段是对应上传文件的名称
	w, err := writer.CreateFormFile("image", path.Base(s))
	if err != nil {
		fmt.Println("创建文件失败: ", err)
	}
	data, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Println("读取图片发生错误: ", err)
	}
	// 把文件内容写入
	w.Write(data)
	writer.Close()
	// 发送一个POST请求
	req, err := http.NewRequest("POST", "https://pic.sogou.com/pic/upload_pic.jsp", &buff)
	if err != nil {
		fmt.Println("创建请求失败: ", err)
	}
	// 设置你需要的Header
	req.Header.Set("Host", "pic.sogou.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Mobile Safari/537.36")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	var client http.Client
	// 执行请求
	resp, err := client.Do(req)
	defer resp.Body.Close()
	// 读取返回内容
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取失败")
	}
	return string(d)
}
