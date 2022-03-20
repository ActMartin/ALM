package main

import (
	"encoding/csv"
	"os"
)

func main(){
	WriteCsv()
}

func WriteCsv() {
	//创建文件
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	//创建一个新的写入文件流
	w := csv.NewWriter(f)
	data := [][]string{
		{"1", "刘备", "23"},
		{"2", "张飞", "23"},
		{"3", "关羽", "23"},
		{"4", "赵云", "23"},
		{"5", "黄忠", "23"},
		{"6", "马超", "23"},
	}
	//写入数据
	w.WriteAll(data)
	w.Flush()
}