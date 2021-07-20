/***********
 * 快速排序
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
	re := quick(&data)

	// 输出结果
	fmt.Println(re)
}

func merge(data *[]int) []int {

	return // 递归处理流程

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
