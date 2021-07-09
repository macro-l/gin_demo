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

	fmt.Println(data)

	// 开始算法
	re := insertion(data)

	// 输出结果
	fmt.Println(re)
}

func insertion(data []int) []int {
	// 需要处理的次数
	// 从第二数开始处理，即需要处理n-1次
	for i := 1; i < len(data); i++ {
		// 每次处理的的循环数
		// 由当前处理点的前一个点开始，--向逐个比较
		for y := i - 1; y >= 0; y-- {
			//fmt.Println(data[i], " => ", data[y])

			if data[i] > data[y] {
				// 无须插入
				if y == i {
					continue
				}
				// 执行插入
				in(data, i, y+1)
				break
			}
			// 首部插入
			if y == 0 {
				in(data, i, 0)
			}
		}
		//fmt.Println(data)
		//fmt.Println()
		//fmt.Println()
	}

	return data
}

func in(data []int, x int, y int) {
	// 将插入操作；影响的区间前移
	for i := x; i > y; i-- {
		data[i], data[i-1] = data[i-1], data[i]
	}
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
