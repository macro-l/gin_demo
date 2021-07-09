/***********
 * 希尔排序
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
	//shell(data)
	re := shell(data)

	// 输出结果
	fmt.Println(re)
	// 验证
	//for i := 1; i < 1000000; i++ {
	//	if 1 != (re[i] - re[i-1]) {
	//		fmt.Println("err")
	//		break
	//	}
	//}
}

func shell(data []int) []int {
	length := len(data)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	fmt.Println(gap)
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := data[i]
			j := i - gap
			for j >= 0 && data[j] > temp {
				data[j+gap] = data[j]
				j -= gap
			}
			data[j+gap] = temp
		}
		gap = gap / 3
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
