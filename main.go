package main

import (
	"fmt"
	"path/filepath"

	"caorushizi.cn/idm1/downloader"
	"caorushizi.cn/idm1/helper"
)

func main() {
	url := "http://xiazai.kugou.com/Corp/kugou7_1099.exe"
	dest := "C:\\Users\\84996\\Desktop\\download\\"

	filename, err := helper.GetFileName(url)
	if err != nil {
		panic(err)
	}

	filepath := filepath.Join(dest, filename)

	supportsRange, err := helper.SupportsRange(url)
	if err != nil {
		panic(err)
	}

	if supportsRange {
		fmt.Println("Supports range")
		size, err := helper.GetFileSize(url)
		if err != nil {
			panic(err)
		}

		parts := []downloader.Part{
			{Start: 0, End: size / 2},
			{Start: size/2 + 1, End: size},
		}
		downloader.DownloadFileParts(url, filepath, parts)
	} else {
		fmt.Println("Does not support range")
		downloader.DownloadFile(url, filepath)
	}
}
