/***
 * 选择排序
 ***/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"time"
)

var (
	//num = 100000
	//rand.Seed(time.Now().UnixNano())
	//sourcesPath     = "./source.txt"
	//sourcesJsonPath = "./source_json.txt"
	sourcesJsonPath = "E:\\go\\gin-demo\\algorithm\\sort\\source_json.txt"
)

func main() {
	startTime := time.Now().UnixNano()
	defer delay(startTime)
	// 获取基础数据
	data := data(sourcesJsonPath)
	fmt.Println(data)

	// 开始算法
	re := selection(data)

	// 输出结果
	fmt.Println(re)
}

func selection(data []int) []int {
	len := len(data)
	// 需要处理的个数
	for i := 0; i < len; i++ {
		// 每次处理需要循环的次数
		// 从当前处理点的下一个点开始，直到长度的最后一个
		for y := i + 1; y < len; y++ {
			minPoint := i
			if data[minPoint] > data[y] {
				data[minPoint], data[y] = data[y], data[minPoint]
			}
		}
		//fmt.Println()
		//fmt.Println()
	}

	return data
}

func data(sourcesJsonPath string) []int {
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
