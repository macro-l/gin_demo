package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// 获取数据
	jsonData, err := ioutil.ReadFile("./source_json.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	var data []int
	json.Unmarshal(jsonData, &data)

	// 排序
	// 处理len-1次,从1开始向后处理
	for i := 1; i < len(data); i++ {
		// 记录下当前处理点
		temp := data[i]
		y := i - 1
		// 由记录点向首部处理
		// 置换条件：比记录点大
		for y >= 0 && data[y] > temp {
			data[y], data[y+1] = data[y+1], data[y]
			y--
		}
		// 停止条件：
		//      1 - 比记录点小的停止处理，并将记录点存储在其后
		//      2 - 当处理到首部是，停止，记录点存储在首部
		data[y+1] = temp
	}

	//fmt.Println(data)
	// 验证
	for i := 1; i < len(data); i++ {
		if 1 != (data[i] - data[i-1]) {
			fmt.Println("err")
			break
		}
	}
	fmt.Println("完成")

}
