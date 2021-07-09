/***
 * 生成测试数据， 写入文件中
 ***/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

var (
	//num = 10
	num = 100000
	//num = 1000000
	rr = rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().UnixNano())
	sourcesPath     = "./source.txt"
	sourcesJsonPath = "./source_json.txt"
)

func main() {
	tdata := []int{}
	data := []byte{}
	for i := 0; i < num; i++ {
		tdata = append(tdata, i)
	}

	for y := 0; y < num; y++ {
		fir := rr.Intn(num)
		//fmt.Printf("%d", fir)

		sec := rr.Intn(num)
		//fmt.Printf("  ==>  %d\n", sec)

		tdata[fir], tdata[sec] = tdata[sec], tdata[fir]
	}

	for x := 0; x < num; x++ {
		tmp := []byte(strconv.Itoa(tdata[x]))

		data = append(data, tmp...)
		data = append(data, []byte{'\n'}...)
	}

	err := ioutil.WriteFile(sourcesPath, data, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(tdata)
	//fmt.Println(json.Marshal(tdata))
	td, _ := json.Marshal(tdata)
	if err := ioutil.WriteFile(sourcesJsonPath, td, 0666); err != nil {
		fmt.Println(err.Error())
	}
}
