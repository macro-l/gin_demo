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
	length := len(data)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if data[min] > data[j] {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
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
