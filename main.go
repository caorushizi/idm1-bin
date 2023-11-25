package main

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(url string, filepath string) error {
	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建文件
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 将响应体数据写入文件
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	url := "http://example.com/file"
	filepath := "/path/to/file"
	err := DownloadFile(url, filepath)
	if err != nil {
		panic(err)
	}
}
