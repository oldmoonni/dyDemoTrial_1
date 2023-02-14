package main

import (
	"time"
)

func main() {
	//去掉文件后缀
	//fullFilename := "test.txt"
	//fmt.Println("filenameOnly =", strings.TrimSuffix(fullFilename, path.Ext(fullFilename)))

	//得到MM-DD类型的日期
	currentTime := time.Now()
	println(currentTime.Format("01-02"))
}
