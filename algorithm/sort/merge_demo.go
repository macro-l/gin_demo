/***********
 * 归并排序
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
	re := mergeSort(data)

	// 输出结果
	fmt.Println(re)
}

func mergeSort(arr []int) []int {
	// 获取总长 （长度记录下来，跟执行性体在同个内存片上，因连续性，回在同个快速缓存内存上，减少内存->高速缓存的读取，增快程序的执行效率）
	length := len(arr)
	// 递归结束节点
	if length < 2 {
		return arr
	}
	// 中间指针
	middle := length / 2
	// 左区数据集合
	left := arr[0:middle]
	// 右区数据集合
	right := arr[middle:]

	// 二分法左右分离，层层递归
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	// 回调结果数据集合
	var result []int
	// 左右区的数据都跑完时终止
	for len(left) != 0 && len(right) != 0 {
		// 左区首位值小于或等于（使数据稳定）右区首位值，取左区首位值
		if left[0] <= right[0] {
			// 取左区首位值
			result = append(result, left[0])
			// 裁剪左区集合，获取[1:len]的子集
			left = left[1:]
		} else { // 反之，右区首位值小于左区首位值，取右区首位值
			// 取右区首位值
			result = append(result, right[0])
			// 裁剪右区集合，获取[1:len]的子集
			right = right[1:]
		}
	}

	// 当右区为空时，循环取左区值
	for len(left) != 0 {
		// 取左区首位值
		result = append(result, left[0])
		// 裁剪左区集合，获取[1:len]的子集
		left = left[1:]
	}

	// 当左区为空时，循环取右区值
	for len(right) != 0 {
		// 取右区首位值
		result = append(result, right[0])
		// 裁剪右区集合，获取[1:len]的子集
		right = right[1:]
	}

	//  回调处理后的数据
	return result
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
