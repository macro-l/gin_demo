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
	re := merge(&data)

	// 输出结果
	fmt.Println(re)
}

func merge(data *[]int) []int {
	dlen := len(*data)

	// 终止递归条件
	if 2 == dlen {
		if (*data)[0] > (*data)[1] {
			(*data)[0], (*data)[1] = (*data)[1], (*data)[0]
		}
		return *data
	}
	if 1 == dlen {
		return *data
	}
	fdata := (*data)[0:(dlen / 2)]
	sdata := (*data)[(dlen / 2):dlen]

	fre := merge(&fdata)
	sre := merge(&sdata)
	flen := len(fre)
	slen := len(sre)
	findex := 0
	sindex := 0

	var re []int
	for findex != flen || sindex != slen {

		if flen == findex {
			re = append(re, sre[sindex])
			sindex++
			continue
		}
		if slen == sindex {
			re = append(re, fre[findex])
			findex++
			continue
		}
		if fre[findex] > sre[sindex] {
			re = append(re, sre[sindex])
			sindex++
		} else {
			re = append(re, fre[findex])
			findex++
		}

	}
	// 递归处理流程

	return re
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
