/***********
 * 插入排序
 ***********/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"time"
)

var (
	sourcesJsonPath = "E:\\go\\gin-demo\\algorithm\\sort\\source_json.txt"
)

func main() {
	// 开始时间
	startTime := time.Now().UnixNano()

	// 计算算法时间花费
	defer delay(startTime)

	// 获取测试数据流
	data := data()
	//fmt.Println(data)

	// 开始算法
	re := shell(data)

	// 输出结果
	fmt.Println(re)
}

func shell(data []int) []int {
	// 增量序列
	len := len(data)
	inc := len / 3

	for inc > 0 {
		fmt.Println(inc)
		// 本次增量的处理
		// 从0开始 当i+inc=len时停止
		for i := inc; i < len; i++ {
			temp := data[i]
			j := i - inc
			for j >= 0 && data[j] > temp {
				data[j+inc] = data[j]
				j -= inc
			}
			data[j+inc] = temp

		}
		// 进入下个增量
		if 6 > inc && 1 < inc {
			inc = 1
		} else {
			inc = inc / 3
		}
	}
	return data
}

func data() []int {
	sources_json, err := ioutil.ReadFile(sourcesJsonPath)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var sources []int
	json.Unmarshal(sources_json, &sources)

	return sources
}

func delay(startTime int64) {
	endTime := time.Now().UnixNano()
	delay := float64(endTime-startTime) / math.Pow(9, 10)
	fmt.Printf("use %.9f \n", delay)
}
